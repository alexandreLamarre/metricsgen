package metricsgen

import (
	"fmt"
	"go/format"
	"log/slog"
	"os"
	"slices"

	"github.com/alexandreLamarre/metricsgen/pkg/metricsgen"
	"github.com/alexandreLamarre/metricsgen/pkg/templates"
	"gopkg.in/yaml.v3"
)

type RunOptions struct {
	MainFile   string
	ExtraFiles []string
	Driver     string
	Package    string
}

var validDrivers = []string{"otel", "prometheus"}

func Run(opts RunOptions) error {
	driver := opts.Driver
	logger := slog.Default().With("driver", driver)
	if !slices.Contains(validDrivers, driver) {
		logger.Error("invalid driver")
		return fmt.Errorf("invalid driver")
	}
	genFile := opts.MainFile
	curPkg := opts.Package
	logger = logger.With("package", curPkg, "metrics-file", genFile)
	logger.Info("reading base configuration")
	data, err := os.ReadFile(genFile)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	cfg := &metricsgen.Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		logger.Error(err.Error())
		return err
	}
	cfg.SetSource(genFile)
	cfg.SetLogger(logger)
	cfg.SetDriver(driver)
	if err := cfg.Sanitize(); err != nil {
		logger.With("stage", "sanitization").Error(err.Error())
		return err
	}
	extraCfgs := []*metricsgen.Config{}
	for _, f := range opts.ExtraFiles {
		logger := logger.With("file", f)
		logger.Info("reading extra configuration")
		data, err := os.ReadFile(f)
		if err != nil {
			logger.Error(err.Error())
			return err
		}

		extraCfg := &metricsgen.Config{}
		if err := yaml.Unmarshal(data, extraCfg); err != nil {
			logger.Error(err.Error())
			return err
		}
		extraCfg.SetSource(f)

		if err := extraCfg.Sanitize(); err != nil {
			logger.With("stage", "sanitization").Error(err.Error())
		}

		extraCfgs = append(extraCfgs, extraCfg)
	}

	if len(extraCfgs) > 0 {
		logger.Info("merging configurations")
		if err := cfg.Merge(extraCfgs...); err != nil {
			logger.Error(err.Error())
			return err
		}
	}
	logger.Info("validating configuration")
	if err := cfg.Validate(); err != nil {
		return err
	}

	var metricsGenRet []byte
	if driver == "otel" {
		metricsGen, err := templates.ExecuteOtelMetrics(templates.GenConfig{
			PackageName: curPkg,
			ImportDefs: []templates.ImportDef{
				{
					Dependency: "context",
				},
				{
					Alias:      "otelmetricsdk",
					Dependency: "go.opentelemetry.io/otel/metric",
				},
				{
					Alias:      "otelattribute",
					Dependency: "go.opentelemetry.io/otel/attribute",
				},
				{
					Alias:      "prommodel",
					Dependency: "github.com/prometheus/common/model",
				},
			},
			Metrics:   cfg.ToMetricsTemplateDefinition(),
			EnumTypes: cfg.ToEnumTemplateDefinition(),
		})
		if err != nil {
			logger.With("stage", "codegen").Error(err.Error())
			return err
		}
		metricsGenRet = metricsGen
	}

	if driver == "prometheus" {
		metricsGen, err := templates.ExecutePrometheusMetrics(
			templates.GenConfig{
				PackageName: curPkg,
				Metrics:     cfg.ToMetricsTemplateDefinition(),
				EnumTypes:   cfg.ToEnumTemplateDefinition(),
				ImportDefs: []templates.ImportDef{
					{
						Dependency: "fmt",
					},
					{
						Dependency: "strings",
					},
					{
						Alias:      "promsdk",
						Dependency: "github.com/prometheus/client_golang/prometheus",
					},
				},
			},
		)
		if err != nil {
			logger.With("stage", "codegen").Error(err.Error())
			return err
		}
		metricsGenRet = metricsGen
	}
	metricsGenRet, err = format.Source(metricsGenRet)
	if err != nil {
		logger.With("stage", "formatting").Error(err.Error())
		return err
	}

	if err := os.WriteFile(
		fmt.Sprintf("metrics_%s_generated.go", driver),
		metricsGenRet,
		0644,
	); err != nil {
		return err
	}

	docsgen, err := templates.ExecuteDocs(cfg.ToDocsTemplateDefinition())
	if err != nil {
		logger.With("stage", "docgen")
		return err
	}

	if err := os.WriteFile("metrics.md", docsgen, 0644); err != nil {
		return err
	}
	return nil
}

func Validate(file string) error {
	genFile := file
	logger := slog.Default()
	logger = logger.With("metrics-file", genFile)

	data, err := os.ReadFile(genFile)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	cfg := &metricsgen.Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		logger.Error(err.Error())
		return err
	}
	cfg.SetSource(genFile)
	cfg.SetLogger(logger)
	if err := cfg.Sanitize(); err != nil {
		logger.With("stage", "sanitization").Error(err.Error())
		return err
	}

	if err := cfg.Validate(); err != nil {
		logger.With("stage", "validation").Error(err.Error())
	}
	logger.Info("configuration is valid")
	return nil
}
