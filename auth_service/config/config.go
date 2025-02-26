package configPkg

import (
	"fmt"
	"log"
	"os"
)

// Config holds database configuration
type Config struct {
	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	DBSSLMode  string `yaml:"db_sslmode"`
	// DBDNS is optional and can be overridden; computed otherwise
	DBDNS string `yaml:"db_dns,omitempty"`
}

// GetDSN returns the PostgreSQL DSN (Data Source Name)
func (c *Config) GetDSN() string {
	if c.DBDNS != "" {
		return c.DBDNS // Use provided DSN if available
	}
	// Compute DSN if not explicitly set
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode)
}

// LoadConfig loads configuration based on the environment
func LoadConfig(env string) (*Config, error) {

	// Load config from environment variables with defaults
	config := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBSSLMode:  getEnv("DB_SSLMODE", "require"), // Default to "require" for RDS
		DBDNS:      getEnv("DB_DNS", ""),
	}

	// Validate required fields
	if config.DBHost == "" || config.DBUser == "" || config.DBPassword == "" || config.DBName == "" {
		return nil, fmt.Errorf("missing required environment variables (DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)")
	}

	// Log the loaded config (masking sensitive data)
	safeDSN := fmt.Sprintf("postgres://%s:***@%s:%s/%s?sslmode=%s",
		config.DBUser, config.DBHost, config.DBPort, config.DBName, config.DBSSLMode)
	log.Printf("Loaded configuration for environment: %s, DSN: %s", env, safeDSN)

	return config, nil
}

// getEnv retrieves an environment variable or returns a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}
