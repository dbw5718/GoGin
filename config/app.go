package config

type App struct {
	Env     string `mapstructe:"env" json:"env" yaml:"env"`
	Port    string `mapstructu:"prt" json:"port" yaml:"port"`
	AppName string `mapstructure:"appname" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}
