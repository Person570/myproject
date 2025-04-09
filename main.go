package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/myproject/handlers" // Абсолютный путь
)

type Config struct {
	ApiUrl   string `json:"api_url"`
	BotToken string `json:"bot_token"`
	ChatId   string `json:"chat_id"`
	LogFile  string `json:"log_file"`
}

func loadConfig() Config {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Ошибка при чтении конфигурации: %v", err)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		log.Fatalf("Ошибка при парсинге конфигурации: %v", err)
	}

	return config
}

func main() {
	config := loadConfig()

	logFile, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Ошибка при настройке логов: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	for {
		products, err := handlers.FetchProducts(config.ApiUrl)
		if err != nil {
			log.Printf("Ошибка при получении данных: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		handlers.AnalyzeProducts(products, config.BotToken, config.ChatId)
		time.Sleep(5 * time.Second)
	}
}
