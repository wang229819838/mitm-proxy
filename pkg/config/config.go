// pkg/config.go

package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	CACert string `json:"ca_cert"`
	CAKey  string `json:"ca_key"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
