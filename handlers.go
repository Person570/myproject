package handlers // Имя пакета должно соответствовать имени папки

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func FetchProducts(apiUrl string) ([]Product, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("ошибка HTTP-запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка HTTP: статус %d", resp.StatusCode)
	}

	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	return products, nil
}

func AnalyzeProducts(products []Product, botToken, chatID string) {
	for _, product := range products {
		log.Printf("Товар: %s, Цена: %.2f\n", product.Name, product.Price)

		msg := fmt.Sprintf("Товар: %s, Цена: %.2f", product.Name, product.Price)
		if err := sendMessage(botToken, chatID, msg); err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
		}
	}
}

func sendMessage(botToken, chatID, message string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	params := url.Values{}
	params.Add("chat_id", chatID)
	params.Add("text", message)

	resp, err := http.PostForm(endpoint, params)
	if err != nil {
		return fmt.Errorf("ошибка при отправке сообщения: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка при отправке сообщения. Код статуса: %d", resp.StatusCode)
	}

	return nil
}
