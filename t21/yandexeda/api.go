package yandexeda

import (
	"encoding/json"
	"log"
)

type YandexEdaAPI struct {
	baseURL string
}

func NewAPI(baseURL string) *YandexEdaAPI {
	return &YandexEdaAPI{baseURL}
}

func (y *YandexEdaAPI) GetFoods() json.RawMessage {
	log.Println("Getting foods from Yandex Eda")

	return []byte(``)
}
