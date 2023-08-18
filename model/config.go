package model

type Config struct {
	HomeAssistant struct {
		IP    string `yaml:"ip"`
		Port  int    `yaml:"port"`
		Path  string `yaml:"path"`
		Token string `yaml:"token"`
	} `yaml:"homeassistant"`
}
