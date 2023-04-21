package config

type Config struct {
	Mysql  Mysql `yaml: "mysql"`
	Logger Mysql `yaml: "logger"`
	System Mysql `yaml: "system"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"post"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       int    `yaml:"prefix"`
	Director     string `yaml:"director"`
	showLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"Port"`
	Env  string `yaml:"env"`
}
