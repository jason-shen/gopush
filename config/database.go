package config

type DatabaseConfig struct {
	Host 	 string
	Name     string
	Password string
	User     string
	Port     string
}

func NewDatabase() *DatabaseConfig {
	return &DatabaseConfig{
		Host:	  getIni("database", "HOST", "localhost"),
		Name:     getIni("database", "DATABASE_NAME", "gopush_dev"),
		User:     getIni("database", "DATABASE_USER", "postgres"),
		Password: getIni("database", "DATABASE_PASSWORD", "postgres"),
		Port:     getIni("database", "DATABASE_PORT", "5432"),
	}
}