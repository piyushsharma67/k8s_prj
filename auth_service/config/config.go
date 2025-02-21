package configPkg

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	DBSSLMode  string `yaml:"db_sslmode"`
	DBDNS      string `yaml:"DB_DNS"`
}

func (config *Config)GetDSN()string{
	return config.DBDNS
}

func Loadconfig(env string) (*Config, error) {
	config := &Config{}

	if env == "local" {
		// Load from config.yaml
		yamlFile, err := os.ReadFile("config.yaml")
		
		if err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}

		var yamlConfig map[string]Config
		
		err = yaml.Unmarshal(yamlFile, &yamlConfig)
		
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal yaml: %w", err)
		}

		if cfg, ok := yamlConfig["local"]; ok {
			config = &cfg
			config.DBDNS=fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName, config.DBSSLMode)
		} else {
			return nil, fmt.Errorf("dev config not found in config.yaml")
		}
	} else {
		config = &Config{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			DBSSLMode:  os.Getenv("DB_SSLMODE"),
			DBDNS:      os.Getenv("DB_DNS"),
		}
	}

	log.Printf("Loaded configuration for environment: %s", env)
	return config, nil
}