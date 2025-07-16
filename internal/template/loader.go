package template

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/*
var TemplatesFS embed.FS

func Load(name string) (*template.Template, error) {
	content, err := TemplatesFS.ReadFile("templates/" + name)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o template '%s': %w", name, err)
	}
	return template.New(name).Parse(string(content))
}
