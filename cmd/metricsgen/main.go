package main

import (
	"os"

	_ "github.com/alexandreLamarre/metricsgen/pkg/logger"

	"github.com/alexandreLamarre/metricsgen"
	"github.com/alexandreLamarre/metricsgen/pkg/version"
	"github.com/spf13/cobra"
)

func BuildGenerateCmd() *cobra.Command {
	var extraFiles []string
	var driver string
	cmd := &cobra.Command{
		Args:    cobra.ExactArgs(1),
		Use:     "metricsgen <filename>",
		Version: version.FriendlyVersion(),
		RunE: func(cmd *cobra.Command, args []string) error {
			var curPkg string
			curPkg = os.Getenv("GOPACKAGE")
			if curPkg == "" {
				curPkg = "metrics"
			}
			return metricsgen.Run(metricsgen.RunOptions{
				MainFile:   args[0],
				ExtraFiles: extraFiles,
				Driver:     driver,
				Package:    curPkg,
			})
		},
	}

	cmd.Flags().StringArrayVarP(&extraFiles, "extra-files", "f", []string{}, "extra metricsgen files to aggregate during generation")
	cmd.Flags().StringVarP(&driver, "driver", "d", "otel", "specifies which sdks to use in code generation. Available: `otel`, `prometheus`")
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate the metricsgen file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return metricsgen.Validate(args[0])
		},
	}
	cmd.AddCommand(validateCmd)

	return cmd
}

func main() {
	cmd := BuildGenerateCmd()
	cmd.Execute()
}
