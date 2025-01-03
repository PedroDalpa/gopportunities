package main

import (
	"github.com/PedroDalpa/gopportunities/config"
	r "github.com/PedroDalpa/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errf("config initialization error: %v", err)
		panic(err)
	}

	r.Init()
}
