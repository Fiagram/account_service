package configs

import (
	"fmt"
	"os"

	"github.com/Fiagram/account_service/configs"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Grpc     Grpc     `yaml:"grpc"`
	Database Database `yaml:"database"`
	Auth     Auth     `yaml:"auth"`
	Log      Log      `yaml:"log"`
}

// Creates a new config instance by reading from a given YAML file.
// If the filePath is empty, it uses the default embedded configuration.
func NewConfig(filePath string) (Config, error) {
	var (
		configBytes = configs.DefaultConfigBytes
		config      = Config{}
		err         error
	)

	if filePath != "" {
		configBytes, err = os.ReadFile(filePath)
		if err != nil {
			return Config{}, fmt.Errorf("Failed to read YAML file: %w", err)
		}
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to unmarshal YAML file: %w", err)
	}

	return config, nil
}
