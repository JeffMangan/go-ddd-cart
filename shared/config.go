package shared

import (
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	AccountID     string
	IsLocal       bool
	LocalEndPoint string
	Region        string
	Key           string
	Secret        string
	LogLevel      string
	LogFormatType string
}

const ConfigPath string = "/home/jeff/Documents/Work/Code/go/go-ddd-cart/shared/config.json"

// NewConfigFromPath loads a config entity from a string
func NewConfigFromPath(path string) (*Config, *CustomError) {
	ext := filepath.Ext(path)
	viper.SetConfigType(ext[1:])
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, NewCustomError(err.Error(), ErrorTypeSystem)
	}
	return &Config{
		AccountID:     viper.GetString("account_id"),
		IsLocal:       viper.GetBool("is_local"),
		Key:           viper.GetString("key"),
		Secret:        viper.GetString("secret"),
		Region:        viper.GetString("region"),
		LocalEndPoint: viper.GetString("local_endpoint"),
		LogLevel:      viper.GetString("log_level"),
		LogFormatType: viper.GetString("log_format_type"),
	}, nil
}
