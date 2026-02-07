package main

import (
	"fmt"
	"github.com/darkbreadmaker/gator/internal/config"
)
func main() {
	cfg := config.Read()
	err := cfg.SetUser("Steve")
	cfg = config.Read()
	fmt.Printf(cfg)
}
