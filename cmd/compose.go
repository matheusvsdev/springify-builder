package cmd

import (
	"github.com/matheusvsdev/springify/cmd/compose"
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Gerenciamento de composição ou estrutura para containers",
}

func init() {
	composeCmd.AddCommand(compose.AddCmd)

	rootCmd.AddCommand(composeCmd)
}
