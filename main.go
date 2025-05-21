package main

import (
	"go/format"
	"log/slog"
	"os"

	_ "github.com/alexandreLamarre/metricsgen/pkg/logger"

	"github.com/alexandreLamarre/metricsgen/pkg/metricsgen"
	"github.com/alexandreLamarre/metricsgen/pkg/templates"
	"github.com/alexandreLamarre/metricsgen/pkg/version"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func BuildGenerateCmd() *cobra.Command {
	var extraFiles []string
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "metricsgen <filename>",
		Version: version.FriendlyVersion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := slog.Default()
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

			metricsgen, err := templates.ExecuteMetrics(templates.GenConfig{
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
				},
				Metrics:   cfg.ToMetricsTemplateDefinition(),
				EnumTypes: cfg.ToEnumTemplateDefinition(),
			})
			if err != nil {
				logger.With("stage", "codegen").Error(err.Error())
				return err
			}

			metricsgen, err = format.Source(metricsgen)
			if err != nil {
				logger.With("stage", "formatting").Error(err.Error())
				return err
			}

			if err := os.WriteFile("metrics_generated.go", metricsgen, 0644); err != nil {
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

	return cmd
}

func main() {
	cmd := BuildGenerateCmd()
	cmd.Execute()
}
