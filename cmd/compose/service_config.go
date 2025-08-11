package compose

import "fmt"

type ServiceConfig struct {
	ImageName           string
	DefaultInternalPort string
	DefaultVolumePath   string
	TemplateFile        string
	EnvVars             func(dbName string) map[string]string
}

func GetServiceConfig(serviceType string) (*ServiceConfig, error) {
	switch serviceType {
	case "postgres":
		return &ServiceConfig{
			ImageName:           "postgres",
			DefaultInternalPort: "5432",
			DefaultVolumePath:   "/var/lib/postgresql/data",
			TemplateFile:        "service_db.yml.tmpl",
			EnvVars: func(dbName string) map[string]string {
				return map[string]string{
					"POSTGRES_DB":       dbName,
					"POSTGRES_USER":     "${POSTGRES_USER}",
					"POSTGRES_PASSWORD": "${POSTGRES_PASSWORD}",
				}
			},
		}, nil
	case "mysql":
		return &ServiceConfig{
			ImageName:           "mysql",
			DefaultInternalPort: "3306",
			DefaultVolumePath:   "/var/lib/mysql",
			TemplateFile:        "service_db.yml.tmpl",
			EnvVars: func(dbName string) map[string]string {
				return map[string]string{
					"MYSQL_DATABASE":      dbName,
					"MYSQL_USER":          "${MYSQL_USER}",
					"MYSQL_PASSWORD":      "${MYSQL_PASSWORD}",
					"MYSQL_ROOT_PASSWORD": "${MYSQL_ROOT_PASSWORD}",
				}
			},
		}, nil
	case "oracle":
		return &ServiceConfig{
			ImageName:           "oracle",
			DefaultInternalPort: "1521",
			DefaultVolumePath:   "/opt/oracle/oradata",
			TemplateFile:        "service_db.yml.tmpl",
			EnvVars: func(dbName string) map[string]string {
				return map[string]string{
					"ORACLE_PWD": "${ORACLE_PWD}",
				}
			},
		}, nil
	default:
		return nil, fmt.Errorf("serviço '%s' não reconhecido", serviceType)
	}
}
