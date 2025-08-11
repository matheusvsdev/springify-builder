package promptCompose

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func AskComposeProjectName() string {
	var composeProjectName string
	survey.AskOne(&survey.Input{Message: "Compose project name:"}, &composeProjectName)
	return composeProjectName
}

func AskServiceName() string {
	var serviceName string
	survey.AskOne(&survey.Input{Message: "Nome do serviço:"}, &serviceName)
	return serviceName
}

func AskImageName() string {
	var imageName string
	survey.AskOne(&survey.Select{
		Message: "Imagem:",
		Options: []string{"postgres", "mysql", "oracle"},
		Default: "postgres",
	}, &imageName)
	return imageName
}

func AskContainerName() string {
	var containerName string
	survey.AskOne(&survey.Input{Message: "Nome do container:"}, &containerName)
	return containerName
}

func AskDbName() string {
	var dbName string
	survey.AskOne(&survey.Input{Message: "Nome do banco:"}, &dbName)
	return dbName
}

func AskPort(defaultPort string) string {
	var port string
	survey.AskOne(&survey.Input{
		Message: fmt.Sprintf("Porta (default %s):", defaultPort),
		Default: defaultPort,
	}, &port)
	return port
}

func AskNetworkName(defaultName string) string {
	var network string
	survey.AskOne(&survey.Input{
		Message: fmt.Sprintf("Nome da rede (default %s):", defaultName),
		Default: defaultName,
	}, &network)
	return network
}
