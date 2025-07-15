package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "springify",
	Short: "CLI para geração de perfis e configuração de ambientes Spring",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
