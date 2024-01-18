package main

import (
	"github.com/diegosilva19/go-api/config"
	"github.com/diegosilva19/go-api/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	logger.Infof("ola mundo")
	err := config.Init()

	if err != nil {
		logger.Errorf("error incial %v", err)
		//panic(err)
		return
	}
	router.Initialize()
}
