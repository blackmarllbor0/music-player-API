package config

type ConfigGetter interface {
	GetFromDB(key string) (string, error)
	GetFromServer(key string) (string, error)
}