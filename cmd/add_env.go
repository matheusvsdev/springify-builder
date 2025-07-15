package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

//go:embed templates/*
var templatesFS embed.FS

type DbConfig struct {
	Dialect string
	Driver  string
	Port    string
}

var dbDefaults = map[string]DbConfig{
	"postgresql": {Dialect: "PostgreSQL", Driver: "org.postgresql.Driver", Port: "5432"},
	"mysql":      {Dialect: "MySQL", Driver: "com.mysql.cj.jdbc.Driver", Port: "3306"},
	"oracle":     {Dialect: "Oracle", Driver: "oracle.jdbc.OracleDriver", Port: "1521"},
	"h2":         {Dialect: "H2", Driver: "org.h2.Driver", Port: ""},
}

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Gerenciamento de perfis de ambiente",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Cria perfil de configuração (test, dev, prod)",
	Run: func(cmd *cobra.Command, args []string) {
		var format string
		survey.AskOne(&survey.Select{
			Message: "Escolha o formato de saída:",
			Options: []string{".properties", ".yml"},
			Default: ".properties",
		}, &format)

		var appName string
		survey.AskOne(&survey.Input{Message: "Nome da aplicação:"}, &appName)

		principalPath := "src/main/resources/application" + format
		baseTemplate := strings.TrimPrefix(format, ".") + ".tmpl"

		baseContent, err := templatesFS.ReadFile("templates/" + baseTemplate)
		if err != nil {
			fmt.Printf("X Erro ao ler template base '%s': %s\n", baseTemplate, err)
			return
		}

		baseTmpl, err := template.New("base").Parse(string(baseContent))
		if err != nil {
			fmt.Println("X Erro ao compilar template base:", err)
			return
		}

		baseFile, err := os.Create(principalPath)
		if err != nil {
			fmt.Println("X Erro ao gerar o arquivo base:", err)
			return
		}
		defer baseFile.Close()

		baseTmpl.Execute(baseFile, map[string]string{"AppName": appName})

		var env string
		survey.AskOne(&survey.Select{
			Message: "Escolha o ambiente:",
			Options: []string{"test", "dev", "prod"},
			Default: "dev",
		}, &env)

		var dbType string
		var dbConfig DbConfig
		var host, port, dbName string

		if env == "test" {
			dbType = "h2"
			dbConfig = dbDefaults["h2"]
		} else {
			survey.AskOne(&survey.Select{
				Message: "Tipo de banco:",
				Options: []string{"postgresql", "mysql", "oracle"},
				Default: "postgresql",
			}, &dbType)

			dbConfig = dbDefaults[strings.ToLower(dbType)]
			if dbConfig.Dialect == "" {
				fmt.Printf("X Banco inválido: %s\n", dbType)
				return
			}

			survey.AskOne(&survey.Input{Message: "Host do banco:"}, &host)
			survey.AskOne(&survey.Input{
				Message: fmt.Sprintf("Porta (default %s):", dbConfig.Port),
				Default: dbConfig.Port,
			}, &port)
			survey.AskOne(&survey.Input{Message: "Nome do banco:"}, &dbName)
		}

		// Define dados pro template
		data := map[string]string{
			"AppName":  appName,
			"Db":       strings.ToLower(dbType),
			"Database": dbConfig.Dialect,
			"Driver":   dbConfig.Driver,
			"Host":     host,
			"Port":     port,
			"DbName":   dbName,
		}

		templatePath := fmt.Sprintf("templates/%s-%s.tmpl", env, strings.TrimPrefix(format, "."))
		outputPath := fmt.Sprintf("src/main/resources/application-%s%s", env, format)

		tmplContent, err := templatesFS.ReadFile(templatePath)
		if err != nil {
			fmt.Printf("X Erro ao ler template '%s': %s\n", templatePath, err)
			return
		}

		tmpl, err := template.New(env).Parse(string(tmplContent))
		if err != nil {
			fmt.Println("X Erro ao compilar template:", err)
			return
		}

		file, err := os.Create(outputPath)
		if err != nil {
			fmt.Println("X Erro ao gerar o arquivo:", err)
			return
		}
		defer file.Close()

		tmpl.Execute(file, data)
		fmt.Printf("✔ Arquivo '%s' gerado com sucesso!\n", outputPath)
	},
}

func init() {
	profileCmd.AddCommand(createCmd)
	rootCmd.AddCommand(profileCmd)
}
