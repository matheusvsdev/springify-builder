package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "syb",
	Short: "CLI para inicialização de projetos Spring Boot com ambientes e serviços integrados",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
