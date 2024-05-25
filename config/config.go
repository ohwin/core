package config

type ServerConfig struct {
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
	Mysql  Mysql  `yaml:"mysql"`
}

type Server struct {
	Port    int    `yaml:"port"`
	Mode    string `yaml:"mode"`
	Version string `yaml:"version"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type Redis struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
