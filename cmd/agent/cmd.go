package main

import (
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/agent"
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	agent := agent.New(cfg)
	agent.Run()
}
