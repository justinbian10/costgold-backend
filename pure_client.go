package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	BASE_URL = "https://public.api.collectpure.com"
)

type PureClient struct {
}

type ProductVariants struct {
	Variants []Offer `json:"variants"`
}

type Offer struct {
	Info OfferInfo `json:"highestOffer"`
}

type OfferInfo struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (p *PureClient) GetProduct(productName string) (*Product, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/product/sku/%s", BASE_URL, productMap[productName])
	req, _ := http.NewRequest("GET", url, nil)

	pureApiKey := os.Getenv("PURE_API_KEY")
	req.Header.Add("x-api-key", pureApiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//res, err := http.Get("http://example.com")
	//body, _ := ioutil.ReadAll(res.Body)
	var test ProductVariants
	err = json.NewDecoder(res.Body).Decode(&test)
	if err != nil {
		return nil, err
	}
	price := test.Variants[0].Info.Price

	return &Product{
		productName,
		price,
	}, nil
}
