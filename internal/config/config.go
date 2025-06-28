package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Connection struct {
	Token       string `yaml:"token,omitempty"`
	Owner       string `yaml:"owner,omitempty"`
	Repo        string `yaml:"repo,omitempty"`
	Uri         string `yaml:"uri,omitempty"`
	DbName      string `yaml:"dbName,omitempty"`
	MinPoolSize int    `yaml:"minPoolSize,omitempty"`
	MaxPoolSize int    `yaml:"maxPoolSize,omitempty"`
	Host        string `yaml:"host,omitempty"`
	Port        int    `yaml:"port,omitempty"`
	Username    string `yaml:"username,omitempty"`
	Password    string `yaml:"password,omitempty"`
}

type Provider struct {
	Name       string     `yaml:"name"`
	Enabled    bool       `yaml:"enabled"`
	Connection Connection `yaml:"connection"`
	Modules    []string   `yaml:"modules"`
}

type Repository struct {
	Providers []Provider `yaml:"providers"`
}

type System struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`  // Read timeout in seconds
	WriteTimeout int    `yaml:"write_timeout"` // Write timeout in seconds
	Name         string `yaml:"name"`
	Version      string `yaml:"version"`
	Description  string `yaml:"description"`
	Timezone     string `yaml:"timezone"`
}

type Config struct {
	System     *System    `yaml:"system"`
	Repository Repository `yaml:"repository"`
}

var conf *Config
var once sync.Once

// LoadConfig loads the application configuration from the specified YAML file path.
// It ensures the configuration is loaded only once using sync.Once. If the configuration
// file cannot be opened or parsed, an error is returned. On success, it returns a pointer
// to the loaded Config struct.
func LoadConfig(configPath string) (*Config, error) {
	var loadErr error
	once.Do(func() {
		file, err := os.Open(configPath)
		if err != nil {
			loadErr = err
			return
		}
		defer file.Close()

		var c Config
		decoder := yaml.NewDecoder(file)
		if err = decoder.Decode(&c); err != nil {
			loadErr = err
			return
		}
		conf = &c
	})
	if loadErr != nil {
		return nil, loadErr
	}
	return conf, nil
}

// GetConfig returns the loaded configuration singleton instance.
// It is a variable to allow for easier testing/mocking.
var GetConfig = func() *Config {
	return conf
}

// GetVersion returns the application version from the loaded configuration.
// If the configuration or version is not available, it returns "unknown".
func GetVersion() string {
	if conf != nil && conf.System != nil {
		return conf.System.Version
	}
	return "unknown"
}
