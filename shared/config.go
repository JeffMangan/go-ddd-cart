package shared

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	AccountID     string `mapstructure:""`
	EndPoint      string `mapstructure:"AWS_DDBLOCAL_ENDPOINT"`
	Region        string `mapstructure:"AWS_DDBLOCAL_REGION"`
	Key           string `mapstructure:""`
	Secret        string `mapstructure:""`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	LogFormatType string `mapstructure:"LOG_FORMAT_TYPE"`
}

func GetConfigPath() string {
	return os.Getenv("CONFIG_FILE_PATH")
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (*Config, *CustomError) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
}

type Config struct {
	AccountID     string
	IsLocal       bool
	EndPoint      string
	Region        string
	Key           string
	Secret        string
	LogLevel      string
	LogFormatType string
}
