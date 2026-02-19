package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.Current_user_name = userName
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error finding home directory: %s", err)
	}
	configPath := fmt.Sprintf("%s/.gatorconfig.json", homeDir)
	jsonData, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Error saving config data: %s", err)
	}
	err = os.WriteFile(configPath, jsonData, 0644)
	return nil
}
