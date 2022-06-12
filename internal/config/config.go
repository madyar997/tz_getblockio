package config

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Client struct {
		ApiKey string `yaml:"apiKey"`
	} `yaml:"client"`
}
