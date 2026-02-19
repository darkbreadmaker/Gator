package main

import (
	"fmt"

	"github.com/darkbreadmaker/internal/config"
)

func main() {
	cfg := config.Read()
	err := cfg.SetUser("Steve")
	if err != nil {
		fmt.Errorf("Error setting user: %s", err)
	}
	cfg = config.Read()
	fmt.Printf("%v", cfg)
}
