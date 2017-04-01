package main



type Config struct {
	port int
}

func (config *Config) ReadConfig(file string) bool {
	config.port = 8080
	
	return true
}

func (config *Config) GetPort() int {
	return config.port;
}
