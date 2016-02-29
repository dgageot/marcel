package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"os"

	"github.com/mitchellh/go-homedir"
)

const path = ".marcel"

// Config describes where to point docker CLI: local daemon, an url or a docker machine.
type Config struct {
	Type     string
	Machine  string
	Url      string
	CertPath string
}

// Save writes the settings to the configuration file.
func Save(config *Config) error {
	buf, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	file, err := file()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, buf, 0644)
}

// Load reads the settings from the configuration file.
func Load() (*Config, error) {
	file, err := file()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return &Config{Type: "local"}, nil
		}

		return nil, err
	}

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config = Config{}
	err = json.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func file() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, path), nil
}
