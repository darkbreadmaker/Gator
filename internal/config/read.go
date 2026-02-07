package config

import (
	"os"
	"fmt"
	"encoding/json"
)

func Read() Config {
	var userConfig Config
	homeDir := os.UserHomeDir()
	jsonData := os.ReadAll(fmt.Sprintf("%s/.gatorconfig.json", homeDir))
	err := json.Unmarshal(jsonData, &userConfig)
	if err != nil {
		fmt.Errorf("Error unmarshalling data: %s", err)
	}
	return userConfig
}
