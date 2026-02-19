package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() Config {
	var userConfig Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Errorf("Error finding home directory: %s", err)
	}
	jsonData, err := os.ReadFile(fmt.Sprintf("%s/.gatorconfig.json", homeDir))
	if err != nil {
		fmt.Errorf("Error reading config file: %s", err)
	}
	err = json.Unmarshal(jsonData, &userConfig)
	if err != nil {
		fmt.Errorf("Error unmarshalling data: %s", err)
	}
	return userConfig
}
