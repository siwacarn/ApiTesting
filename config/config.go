package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct{
	Dialect string
	Username string
	Password string
	Hostname string
	Name string
	Charset string
}

func GetConfig() *Config{
	return &Config{
		DB: &DBConfig{
			Dialect: "mysql",
			Username: "root",
			Password: "root",
			Hostname: "localhost:3306",
			Name: "farm",
			Charset: "utf8",
		}
	}
}