package config

type DbConfig struct {
	Dialect string
	Driver  string
	Port    string
}

var DbDefaults = map[string]DbConfig{
	"postgresql": {Dialect: "PostgreSQL", Driver: "org.postgresql.Driver", Port: "5432"},
	"mysql":      {Dialect: "MySQL", Driver: "com.mysql.cj.jdbc.Driver", Port: "3306"},
	"oracle":     {Dialect: "Oracle", Driver: "oracle.jdbc.OracleDriver", Port: "1521"},
	"h2":         {Dialect: "H2", Driver: "org.h2.Driver", Port: ""},
}
