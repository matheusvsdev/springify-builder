package template

import (
	"bytes"
	"text/template"
)

// Render processa um template com os dados e retorna como string
func Render(tmpl *template.Template, data any) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
