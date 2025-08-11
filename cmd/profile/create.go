package profile

import (
	"fmt"
	"strings"

	"github.com/matheusvsdev/springify/internal/config"
	"github.com/matheusvsdev/springify/internal/model"
	promptProfile "github.com/matheusvsdev/springify/internal/prompt/profile"
	"github.com/matheusvsdev/springify/internal/service"
	"github.com/matheusvsdev/springify/internal/template"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Cria perfil de configuração (test, dev, prod)",
	Run: func(cmd *cobra.Command, args []string) {
		format := promptProfile.AskFormat()
		appName := promptProfile.AskAppName()

		baseTemplate := strings.TrimPrefix(format, ".") + ".tmpl"
		baseTmpl, err := template.Load(baseTemplate)
		if err != nil {
			fmt.Printf("X Erro ao carregar template base: %s\n", err)
			return
		}

		err = service.Generate("src/main/resources/application"+format, baseTmpl, model.BaseData{AppName: appName})
		if err != nil {
			fmt.Println("X Erro ao gerar arquivo base:", err)
			return
		}

		env := promptProfile.AskEnv()

		var dbType string
		var dbConfig config.DbConfig
		var host, port, dbName string

		if env == "test" {
			dbType = "h2"
			dbConfig = config.DbDefaults["h2"]
		} else {
			dbType = promptProfile.AskDbType()
			dbConfig = config.DbDefaults[strings.ToLower(dbType)]

			if dbConfig.Dialect == "" {
				fmt.Println("X Banco inválido:", dbType)
				return
			}

			host = promptProfile.AskHost()
			port = promptProfile.AskPort(dbConfig.Port)
			dbName = promptProfile.AskDbName()
		}

		envTemplate := fmt.Sprintf("%s-%s.tmpl", env, strings.TrimPrefix(format, "."))
		tmpl, err := template.Load(envTemplate)
		if err != nil {
			fmt.Println("X Erro ao carregar template do ambiente:", err)
			return
		}

		// Define dados pro template
		profileData := model.ProfileData{
			AppName:  appName,
			Db:       strings.ToLower(dbType),
			Database: dbConfig.Dialect,
			Driver:   dbConfig.Driver,
			Host:     host,
			Port:     port,
			DbName:   dbName,
		}

		outputPath := fmt.Sprintf("src/main/resources/application-%s%s", env, format)
		err = service.Generate(outputPath, tmpl, profileData)
		if err != nil {
			fmt.Println("X Erro ao gerar o arquivo:", err)
			return
		}

		fmt.Printf("✔ Arquivo '%s' gerado com sucesso!\n", outputPath)
	},
}
