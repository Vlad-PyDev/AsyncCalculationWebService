package main

import (
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/orchestrator"
)

func main() {
	app := orchestrator.New()

	app.Run()
}
