package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	TimeAddition        time.Duration
	TimeSubtraction     time.Duration
	TimeMultiplication  time.Duration
	TimeDivision        time.Duration
	ComputingPower      int
	OrchestratorAddress string
}

func LoadConfig() Config {
	config := Config{
		TimeAddition:        2000 * time.Millisecond,
		TimeSubtraction:     2000 * time.Millisecond,
		TimeMultiplication:  3000 * time.Millisecond,
		TimeDivision:        3000 * time.Millisecond,
		ComputingPower:      3,
		OrchestratorAddress: "localhost:5000",
	}

	envFile, err := os.Open(".env")
	if err != nil {
		log.Println("Unable to open .env file. Applying default configuration.")
		return config
	}
	defer envFile.Close()

	fileScanner := bufio.NewScanner(envFile)
	for fileScanner.Scan() {
		configLine := fileScanner.Text()
		if len(configLine) == 0 || strings.HasPrefix(configLine, "#") {
			continue
		}

		configParts := strings.SplitN(configLine, "=", 2)
		if len(configParts) != 2 {
			continue
		}
		configKey := strings.TrimSpace(configParts[0])
		configValue := strings.TrimSpace(configParts[1])

		switch configKey {
		case "TIME_ADDITION_MS":
			if parsedValue, err := strconv.Atoi(configValue); err == nil && parsedValue > 0 {
				config.TimeAddition = time.Duration(parsedValue) * time.Millisecond
			}
		case "TIME_SUBTRACTION_MS":
			if parsedValue, err := strconv.Atoi(configValue); err == nil && parsedValue > 0 {
				config.TimeSubtraction = time.Duration(parsedValue) * time.Millisecond
			}
		case "TIME_MULTIPLICATIONS_MS":
			if parsedValue, err := strconv.Atoi(configValue); err == nil && parsedValue > 0 {
				config.TimeMultiplication = time.Duration(parsedValue) * time.Millisecond
			}
		case "TIME_DIVISIONS_MS":
			if parsedValue, err := strconv.Atoi(configValue); err == nil && parsedValue > 0 {
				config.TimeDivision = time.Duration(parsedValue) * time.Millisecond
			}
		case "COMPUTING_POWER":
			if parsedValue, err := strconv.Atoi(configValue); err == nil && parsedValue > 0 {
				config.ComputingPower = parsedValue
			}
		case "ORCHESTRATOR_ADDRESS":
			config.OrchestratorAddress = configValue
		}
	}

	if scanErr := fileScanner.Err(); scanErr != nil {
		log.Println("Failed to read .env file:", scanErr)
	}

	return config
}
