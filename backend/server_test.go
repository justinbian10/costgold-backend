package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubScraper struct {
	prices map[string]float64
}

func (s *StubScraper) GetProduct(productName string) Product {
	price := s.prices[productName]
	return Product{
		FullName: productName,
		Price:    price,
	}
}

func TestGETPrice(t *testing.T) {
	scraper := StubScraper{
		map[string]float64{
			"one-oz-fortuna": 20,
			"one-oz-koi":     10,
		},
	}
	server := NewCostgoldServer(&scraper)
	t.Run("returns one oz fortuna's price", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/price/one-oz-fortuna", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got Product
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q int Product, '%v'", response.Body, err)
		}

		want := Product{
			FullName: "one-oz-fortuna",
			Price:    20,
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
