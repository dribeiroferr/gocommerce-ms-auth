package main

import commons "github.com/dribeiroferr/gocommerce-ms-auth/common"

func main() {
	config, err := LoadConfig()
	if err != nil {
		commons.Log.Fatalf("Failed to load configuration: %v", err)
	}
	commons.Log.Infof("envs loaded: %+v", config)
}