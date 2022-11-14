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
	//os.Setenv("CONFIG_FILE_PATH", "....") pull from env variable as seen below
	return os.Getenv("CONFIG_FILE_PATH")
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config *Config, customError *CustomError) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		customError = NewCustomError(err.Error(), ErrorTypeSystem)
		return
	} else {
		err = viper.Unmarshal(&config)
		if err != nil {
			customError = NewCustomError(err.Error(), ErrorTypeSystem)
			return
		}
	}
	return
}
