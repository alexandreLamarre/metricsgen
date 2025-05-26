package main

import (
	"fmt"
	"go/format"
	"log/slog"
	"os"
	"slices"

	_ "github.com/alexandreLamarre/metricsgen/pkg/logger"

	"github.com/alexandreLamarre/metricsgen/pkg/metricsgen"
	"github.com/alexandreLamarre/metricsgen/pkg/templates"
	"github.com/alexandreLamarre/metricsgen/pkg/version"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var validDrivers = []string{"otel", "prometheus"}

func BuildGenerateCmd() *cobra.Command {
	var extraFiles []string
	var driver string
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "metricsgen <filename>",
		Version: version.FriendlyVersion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := slog.Default().With("driver", driver)
			if !slices.Contains(validDrivers, driver) {
				logger.Error("invalid driver")
				return fmt.Errorf("invalid driver")
			}
			genFile := args[0]
			var curPkg string
			curPkg = os.Getenv("GOPACKAGE")
			if curPkg == "" {
				curPkg = "metrics"
			}
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
			for _, f := range extraFiles {
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
		},
	}

	cmd.Flags().StringArrayVarP(&extraFiles, "extra-files", "f", []string{}, "extra metricsgen files to aggregate during generation")
	cmd.Flags().StringVarP(&driver, "driver", "d", "otel", "specifies which sdks to use in code generation. Available: `otel`, `prometheus`")
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate the metricsgen file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			genFile := args[0]
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

		},
	}
	cmd.AddCommand(validateCmd)

	return cmd
}

func main() {
	cmd := BuildGenerateCmd()
	cmd.Execute()
}
