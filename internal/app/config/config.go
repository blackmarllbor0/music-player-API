package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func NewViperConfig() ConfigGetter {
	return &Config{viper.New()}
}

func (c Config) LoadConfig(pathToConfigFile string) error {
	filename := filepath.Base(pathToConfigFile)

	c.SetConfigName(filepath.Base(filename))
	c.SetConfigType(strings.Split(filename, ".")[1])
	c.AddConfigPath(filepath.Dir(pathToConfigFile))

	if err := c.ReadInConfig(); err != nil {
		return fmt.Errorf("error while load config: %v\n", err)
	}

	return nil
}

func (c Config) getValueByKeys(key string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", fmt.Errorf("key cannot be an empty\n")
	}

	value := c.GetString(key)
	if strings.TrimSpace(value) == "" {
		return "", fmt.Errorf(
			"an empty value returned, perhaps it's not defined or is missing\nKey: %s\n",
			key,
		)
	}

	return value, nil
}

func (c Config) GetFromDB(key string) (string, error) {
	return c.getValueByKeys(fmt.Sprintf("DB.%s", key))
}

func (c Config) GetFromServer(key string) (string, error) {
	return c.getValueByKeys(fmt.Sprintf("Server.%s", key))
}
