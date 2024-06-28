package main

import (
	commons "github.com/dribeiroferr/gocommerce-ms-auth/common"
	config "github.com/dribeiroferr/gocommerce-ms-auth/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		commons.Log.Fatalf("Failed to load configuration: %v", err)
	}
	commons.Log.Infof("Envs loaded: %+v", config)
}
