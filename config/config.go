package config

type ServerConfig struct {
	Server Server `yaml:"server"`
	Redis  Redis  `yaml:"redis"`
	Mysql  Mysql  `yaml:"mysql"`
	Casbin bool   `yaml:"casbin"`
}

type Server struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Mode    string `yaml:"mode"`
	Port    string `yaml:"port"`
}

func (s *Server) GetPort() string {
	if len(s.Port) == 0 {
		return ":8080"
	}
	return ":" + s.Port
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
