package cmd

import (
	"github.com/karchx/envtoyaml/pkg"
	"github.com/spf13/cobra"
)

var (
	filePath string
	parseCmd = &cobra.Command{
		Use:   "parse",
		Short: "A brief description of your command",
		Long:  ``,
		Run:   parseFileEnvs,
	}
)

func parseFileEnvs(cmd *cobra.Command, args []string) {
	data := pkg.ReadFile(filePath)
	pkg.CreateFile("yaml", data)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "File env to parse")
	rootCmd.AddCommand(parseCmd)
}
