package config

type ConfigReader interface {
	SetConfigName(name string)
	SetConfigType(t string)
	AddConfigPath(path string)
	ReadInConfig() error
	GetString(key string) string
}