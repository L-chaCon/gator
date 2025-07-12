package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	config, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	configFilePath, err := getFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configFilePath, config, 0644)
	if err != nil {
		return nil
	}
	return nil
}

func Read() (Config, error) {
	cfg := Config{}

	configFilePath, err := getFilePath()
	if err != nil {

	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func getFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/" + configFileName, nil
}
