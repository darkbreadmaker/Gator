package config

import (
	"os"
	"encoding/json"
	"fmt"
)
type Config struct {
	db_url string
	current_user_name string
}

func (c *Config) SetUser(userName string) error{
	c.current_user_name = userName
	homeDir := os.UserHomeDir()
	configPath := fmt.Sprintf("%s/.gatorconfig.json")
	jsonData, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Error saving config data: %s", err)
	}
	err := os.WriteFile(configPath, jsonData)
	return nil
}
