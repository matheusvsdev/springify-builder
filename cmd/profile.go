package cmd

import (
	"github.com/matheusvsdev/springify/cmd/profile"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Gerenciamento de perfils de ambiente",
}

func init() {
	profileCmd.AddCommand(profile.CreateCmd)
	rootCmd.AddCommand(profileCmd)
}
