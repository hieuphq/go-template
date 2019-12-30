package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	jwtKeyPrefix    = "JWT_KEY_"
	jwtSecretPrefix = "JWT_SECRET_"
)

// Config contain configuration of db for migrator
// config var < env < command flag
type Config struct {
	ServiceName     string
	BaseURL         string
	Port            string
	Env             string
	AllowedOrigins  string
	JWTSecret       string
	AccessTokenTTL  int64
	RefreshTokenTTL int64
	DBHost          string
	DBPort          string
	DBUser          string
	DBName          string
	DBPass          string
	DBSSLMode       string
}

// JWTConfig save jwt config key and secret
type JWTConfig struct {
	Key    string
	Secret []byte
	Source string
}

// GetCORS in config
func (c *Config) GetCORS() []string {
	cors := strings.Split(c.AllowedOrigins, ";")
	rs := []string{}
	for idx := range cors {
		itm := cors[idx]
		if strings.TrimSpace(itm) != "" {
			rs = append(rs, itm)
		}
	}
	return rs
}

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

// generateConfigFromViper generate config from viper data
func generateConfigFromViper(v *viper.Viper) Config {

	return Config{
		Port:        v.GetString("PORT"),
		BaseURL:     v.GetString("BASE_URL"),
		ServiceName: v.GetString("SERVICE_NAME"),
		Env:         v.GetString("ENV"),

		AllowedOrigins:  v.GetString("ALLOWED_ORIGINS"),
		AccessTokenTTL:  v.GetInt64("ACCESS_TOKEN_TTL"),
		RefreshTokenTTL: v.GetInt64("REFRESH_TOKEN_TTL"),
		JWTSecret:       v.GetString("JWT_SECRET"),

		DBHost:    v.GetString("DB_HOST"),
		DBPort:    v.GetString("DB_PORT"),
		DBUser:    v.GetString("DB_USER"),
		DBName:    v.GetString("DB_NAME"),
		DBPass:    v.GetString("DB_PASS"),
		DBSSLMode: v.GetString("DB_SSL_MODE"),
	}
}

// DefaultConfigLoaders is default loader list
func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("ENV", "dev")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}
