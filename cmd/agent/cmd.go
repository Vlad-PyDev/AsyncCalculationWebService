package main

import (
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/agent"
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/config"
)

func main() {
	config_ := config.LoadConfig()

	agent := agent.New(config_)
	agent.Run()
}
