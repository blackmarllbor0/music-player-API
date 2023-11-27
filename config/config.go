package config

import (
	"fmt"
	"strings"
)

type Config struct {
	reader ConfigReader
	ConfigGetter
}

func NewConfig(configReader ConfigReader) *Config {
	return &Config{reader: configReader}
}

func (c Config) LoadConfig(
	pathToConfigFile,
	configFileType,
	configFileName string,
) error {
	c.reader.SetConfigName(configFileName)
	c.reader.SetConfigType(configFileType)
	c.reader.AddConfigPath(pathToConfigFile)

	if err := c.reader.ReadInConfig(); err != nil {
		return fmt.Errorf("error while load config: %v\n", err)
	}

	return nil
}

func (c Config) GetValueByKeys(key string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", fmt.Errorf("key cannot be an empty\n")
	}
	
	value := c.reader.GetString(key)
	if strings.TrimSpace(value) == "" {
		return "", fmt.Errorf(
			"an empty value returned, perhaps it's not defined or is missing\nKey: %s\n",
			key,
		)
	}

	return value, nil
}

func (c Config) GetFromDB(key string) (string, error) {
	return c.GetValueByKeys(fmt.Sprintf("DB.%s", key))
}

func (c Config) GetFromServer(key string) (string, error) {
	return c.GetValueByKeys(fmt.Sprintf("Server.%s", key))
}