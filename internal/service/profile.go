package service

import (
	"fmt"
	"os"
	"text/template"
)

func Generate(path string, tmpl *template.Template, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo '%s': %w", path, err)
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
