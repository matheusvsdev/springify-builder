package promptProfile

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func AskFormat() string {
	var format string
	survey.AskOne(&survey.Select{
		Message: "Formato:",
		Options: []string{".properties", ".yml"},
		Default: ".properties",
	}, &format)
	return format
}

func AskAppName() string {
	var appName string
	survey.AskOne(&survey.Input{Message: "Nome da aplicação:"}, &appName)
	return appName
}

func AskEnv() string {
	var env string
	survey.AskOne(&survey.Select{
		Message: "Ambiente:",
		Options: []string{"test", "dev", "prod"},
		Default: "dev",
	}, &env)
	return env
}

func AskDbType() string {
	var dbType string
	survey.AskOne(&survey.Select{
		Message: "Tipo de banco:",
		Options: []string{"postgresql", "mysql", "oracle"},
		Default: "postgresql",
	}, &dbType)
	return dbType
}

func AskHost() string {
	var host string
	survey.AskOne(&survey.Input{Message: "Host do banco:"}, &host)
	return host
}

func AskPort(defaultPort string) string {
	var port string
	survey.AskOne(&survey.Input{
		Message: fmt.Sprintf("Porta (default %s):", defaultPort),
		Default: defaultPort,
	}, &port)
	return port
}

func AskDbName() string {
	var dbName string
	survey.AskOne(&survey.Input{Message: "Nome do banco:"}, &dbName)
	return dbName
}
