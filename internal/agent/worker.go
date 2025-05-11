package agent

import (
	"log"
	"strconv"
	"time"

	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/config"
	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/models"
)

func calculate(arg1, arg2, operator string, cfg config.Config) (float64, string) {
	firstValue, _ := strconv.ParseFloat(arg1, 64)
	secondValue, _ := strconv.ParseFloat(arg2, 64)

	switch operator {
	case "*":
		time.Sleep(cfg.TimeMultiplication)
		return firstValue * secondValue, ""
	case "/":
		time.Sleep(cfg.TimeDivision)
		if secondValue == 0 {
			return 0, "division by zero"
		}
		return firstValue / secondValue, ""
	case "+":
		time.Sleep(cfg.TimeAddition)
		return firstValue + secondValue, ""
	case "-":
		time.Sleep(cfg.TimeSubtraction)
		return firstValue - secondValue, ""
	default:
		return 0, ""
	}
}

func worker(cfg config.Config) {
	for task := range tasksCh {
		log.Printf("Worker received task with ID %v", task.ID)
		calcResult, calcError := calculate(task.Arg1, task.Arg2, task.Type, cfg)

		taskResult := &models.Result{ID: task.ID, Result: calcResult, Error: calcError}
		resultsCh <- taskResult
		log.Printf("Worker sent result %v for task ID %v", calcResult, task.ID)
	}
}
