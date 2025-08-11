package compose

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matheusvsdev/springify/internal/model"
	promptCompose "github.com/matheusvsdev/springify/internal/prompt/compose"
	"github.com/matheusvsdev/springify/internal/service"
	"github.com/matheusvsdev/springify/internal/template"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [serviço]",
	Short: "Adiciona serviços ao docker-compose",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		serviceType := args[0]

		// 🔹 Caminho do docker-compose
		outputPath := filepath.Join("docker", "docker-compose.yml")

		// 🔸 Verifica se o arquivo existe
		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			fmt.Println("⚠ docker-compose.yml não encontrado. Criando base...")

			// Pergunta nome do projeto e da rede
			composeProjectName := promptCompose.AskComposeProjectName()
			network := promptCompose.AskNetworkName(composeProjectName)

			// Carrega template base
			tmpl, err := template.Load("compose-base.yml.tmpl")
			if err != nil {
				fmt.Println("X Erro ao carregar template base:", err)
				return
			}

			// Cria pasta docker se não existir
			err = os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
			if err != nil {
				fmt.Println("X Erro ao criar diretório:", err)
				return
			}

			// Dados para o template base
			composeData := model.ComposeBaseData{
				ComposeProjectName: composeProjectName,
				NetworkName:        network,
			}

			// Gera o arquivo base
			err = service.Generate(outputPath, tmpl, composeData)
			if err != nil {
				fmt.Println("X Erro ao gerar docker-compose base:", err)
				return
			}

			fmt.Println("✔ docker-compose base criado com sucesso!")
		}

		// 🔹 Agora segue o fluxo normal
		config, err := GetServiceConfig(serviceType)
		if err != nil {
			fmt.Println("X Serviço não suportado:", err)
			return
		}

		serviceName := promptCompose.AskServiceName()
		imageName := config.ImageName
		containerName := promptCompose.AskContainerName()
		dbName := promptCompose.AskDbName()
		port := promptCompose.AskPort(config.DefaultInternalPort)

		tmpl, err := template.Load(config.TemplateFile)
		if err != nil {
			fmt.Println("X Erro ao carregar template:", err)
			return
		}

		network, err := service.GetComposeNetworkName(outputPath)
		if err != nil {
			fmt.Println("X não foi possível encontrar rede definida:", err)
			return
		}

		composeData := model.ComposeData{
			ServiceName:   serviceName,
			ImageName:     imageName,
			ContainerName: containerName,
			DbName:        dbName,
			Port:          port,
			InternalPort:  config.DefaultInternalPort,
			VolumePath:    config.DefaultVolumePath,
			NetworkName:   network,
			EnvVars:       config.EnvVars(dbName),
		}

		serviceYaml, err := template.Render(tmpl, composeData)
		if err != nil {
			fmt.Println("X erro ao renderizar serviço:", err)
			return
		}

		err = service.AppendServiceToCompose(outputPath, serviceYaml)
		if err != nil {
			fmt.Println("X erro ao adicionar serviço ao compose:", err)
			return
		}

		fmt.Printf("✔ Serviço '%s' adicionado com sucesso!\n", serviceType)
	},
}

func EnsureComposeBaseExists(outputPath string) error {
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		fmt.Println("⚠ docker-compose.yml não encontrado. Criando base...")

		composeProjectName := promptCompose.AskComposeProjectName()
		network := promptCompose.AskNetworkName(composeProjectName)

		tmpl, err := template.Load("compose-base.yml.tmpl")
		if err != nil {
			return fmt.Errorf("erro ao carregar template base: %w", err)
		}

		err = os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
		if err != nil {
			return fmt.Errorf("erro ao criar diretório: %w", err)
		}

		composeData := model.ComposeBaseData{
			ComposeProjectName: composeProjectName,
			NetworkName:        network,
		}

		err = service.Generate(outputPath, tmpl, composeData)
		if err != nil {
			return fmt.Errorf("erro ao gerar arquivo base: %w", err)
		}

		fmt.Println("✔ docker-compose base criado com sucesso!")
	}
	return nil
}
