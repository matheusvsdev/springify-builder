package service

import (
	"fmt"
	"os"
	"strings"
)

func AppendServiceToCompose(composePath string, serviceYAML string) error {
	content, err := os.ReadFile(composePath)
	if err != nil {
		return fmt.Errorf("falha ao ler compose: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	var result []string
	foundServices := false
	inserted := false

	for i, line := range lines {
		result = append(result, line)

		if strings.HasPrefix(line, "services:") && !inserted {
			serviceIndented := indentYamlBlock(serviceYAML, 2)
			result = append(result, serviceIndented)
			inserted = true
		}

		if foundServices && !inserted {
			if i+1 < len(lines) && !strings.HasPrefix(lines[i+1], "  ") {
				serviceIndented := indentYamlBlock(serviceYAML, 2)
				result = append(result, serviceIndented)
				inserted = true
			}
		}
	}

	if !inserted && foundServices {
		serviceIndented := indentYamlBlock(serviceYAML, 2)
		result = append(result, serviceIndented)
	}

	output := strings.Join(result, "\n")
	err = os.WriteFile(composePath, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("falha ao salvar compose atualizado: %w", err)
	}

	return nil
}

func indentYamlBlock(block string, spaces int) string {
	indent := strings.Repeat(" ", spaces)
	lines := strings.Split(block, "\n")
	var result []string
	skipping := true

	for _, line := range lines {
		if skipping && strings.TrimSpace(line) == "" {
			continue
		}
		skipping = false
		if line == "" {
			result = append(result, "")
		} else {
			result = append(result, indent+line)
		}
	}

	for len(result) > 0 && strings.TrimSpace(result[len(result)-1]) == "" {
		result = result[:len(result)-1]
	}

	return strings.Join(result, "\n")
}
