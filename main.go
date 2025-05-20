package main

import (
	"go/format"
	"os"

	"github.com/alexandreLamarre/metricsgen/pkg/metricsgen"
	"github.com/alexandreLamarre/metricsgen/pkg/templates"
	"github.com/alexandreLamarre/metricsgen/pkg/version"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func BuildGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "metricsgen <filename>",
		Version: version.FriendlyVersion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			genFile := args[0]
			var curPkg string
			curPkg = os.Getenv("GOPACKAGE")
			if curPkg == "" {
				curPkg = "metrics"
			}

			data, err := os.ReadFile(genFile)
			if err != nil {
				return err
			}
			cfg := &metricsgen.Config{}
			if err := yaml.Unmarshal(data, cfg); err != nil {
				return err
			}
			if err := cfg.Sanitize(); err != nil {
				return err
			}
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
				return err
			}

			metricsgen, err = format.Source(metricsgen)
			if err != nil {
				return err
			}

			if err := os.WriteFile("metrics_generated.go", metricsgen, 0644); err != nil {
				return err
			}

			docsgen, err := templates.ExecuteDocs(cfg.ToDocsTemplateDefinition())
			if err != nil {
				return err
			}

			if err := os.WriteFile("metrics.md", docsgen, 0644); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

func main() {
	cmd := BuildGenerateCmd()
	cmd.Execute()
}
