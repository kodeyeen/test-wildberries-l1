package deliveryclub

import (
	"encoding/json"
	"log"
)

type DeliveryClubAPI struct {
	baseURL string
}

func NewAPI(baseURL string) *DeliveryClubAPI {
	return &DeliveryClubAPI{baseURL}
}

func (d *DeliveryClubAPI) FetchProducts() json.RawMessage {
	log.Println("Fetching products from Delivery Club")

	return []byte(``)
}
