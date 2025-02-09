package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Product struct {
	FullName string
	Price    float64
}

type Scraper interface {
	GetProduct(productName string) (*Product, error)
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/price/:product", app.getPriceHandler)

	return app.enableCORS(router)
}

func (app *application) getPriceHandler(w http.ResponseWriter, r *http.Request) {
	productName := httprouter.ParamsFromContext(r.Context()).ByName("product")
	product, err := app.scraper.GetProduct(productName)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(product)

	w.Header().Set("content-type", "application/json")
}

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")

		origin := r.Header.Get("Origin")
		if origin != "" {
			for _, trusted := range app.config.cors.trustedOrigins {
				if origin == trusted {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}
