package config

type Server struct {
	Port int `mapstructure:"Port"`
}

type DB struct {
	User string `mapstructure:"User"`
	Pwd string `mapstructure:"Pwd"`
	DB string `mapstructure:"DB"`
	Port int `mapstructure:"Port"`
}

type App struct {
	Server Server `mapstructure:"Server"`
	DB DB `mapstructure:"DB"`
}