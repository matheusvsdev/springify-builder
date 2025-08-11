package service

import (
	"errors"
	"os"
	"strings"
)

// GetComposeNetworkName lê o docker-compose.yml e retorna o nome da primeira rede definida
func GetComposeNetworkName(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("falha ao ler docker-compose.yml")
	}

	lines := strings.Split(string(content), "\n")
	start := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "networks:" {
			start = true
			continue
		}
		if start {
			// Pega a primeira linha após "networks:" com conteúdo
			if len(trimmed) > 0 && !strings.HasPrefix(trimmed, "#") && strings.Contains(trimmed, ":") {
				return strings.Split(trimmed, ":")[0], nil
			}
		}
	}

	return "", errors.New("nenhuma rede encontrada no docker-compose")
}
