package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const configFilename = ".marcel"

type Config struct {
	Type      string
	Machine   string
	Url       string
	TlsVerify bool
	CertPath  string
}

func Save(config *Config) error {
	buf, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	file, err := configFile()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, buf, 0644)
}

func Load() (*Config, error) {
	file, err := configFile()
	if err != nil {
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

func configFile() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFilename), nil
}
