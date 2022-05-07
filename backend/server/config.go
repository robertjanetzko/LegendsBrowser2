package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	path           string
	LastPath       string
	LastFile       string
	DebugTemplates bool `json:"DebugTemplates,omitempty"`
	DebugJSON      bool `json:"DebugJSON,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	var err error

	if path == "" {
		path, err = configPath()
		if err != nil {
			return nil, err
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("OPEN", err)
		if os.IsNotExist(err) {
			fmt.Println("EX", err)
			home, err := os.UserHomeDir()
			if err != nil {
				return nil, err
			}

			return &Config{LastPath: home, path: path}, nil
		} else {
			return nil, err
		}
	}

	c := &Config{}
	json.Unmarshal(data, c)
	c.path = path
	return c, nil
}

func (c *Config) Save() error {
	var err error

	path := c.path

	if path == "" {
		path, err = configPath()
		if err != nil {
			return err
		}
	}

	file, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, file, 0644)
}

func configPath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, ".legendsbrowser")
	os.MkdirAll(path, os.ModePerm)

	return filepath.Join(path, "config.json"), nil
}
