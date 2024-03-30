package cmd

import (
	"github.com/karchx/envtoyaml/pkg/fs"
	"github.com/spf13/cobra"
)

var (
	filePath  string
	extension string
	parseCmd  = &cobra.Command{
		Use:   "parse",
		Short: "A brief description of your command",
		Long:  ``,
		Run:   parseFileEnvs,
	}
)

func parseFileEnvs(cmd *cobra.Command, args []string) {
	data := fs.ReadFile(filePath)
	fs.CreateFile(extension, data)
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "File env to parse")
	rootCmd.PersistentFlags().StringVarP(&extension, "extension", "e", "yaml", "Extension to Parse")
	rootCmd.AddCommand(parseCmd)
}
