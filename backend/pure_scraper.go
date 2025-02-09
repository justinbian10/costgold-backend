package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const (
	pureURL = "https://www.collectpure.com/marketplace/product/"
)

var productMap = map[string]string{
	"one-oz-fortuna":         "1-oz-pamp-fortuna-gold-bar-9999-fine-in-assay000023",
	"one-oz-koi":             "1-oz-gold-bar-pamp-suisse-good-luck-koi-fish0080",
	"one-oz-eagle-coin-25":   "2025-american-gold-eagle-1-oz-50-coin0111",
	"one-oz-maple-leaf":      "canadian-gold-maple-leaf-1-oz-2024000208",
	"one-oz-buffalo-coin-25": "2025-american-gold-buffalo-1-oz-50-9999-fine-24k-gold-coin0110",
	"one-oz-rand":            "1-oz-rand-refinery-gold-bar-9999-fine-sealed-in-assay-card000108",
	"one-oz-any":             "random-brand-1-oz-gold-bar-9999-fine-in-card000087",
	"fifty-gram-fortuna":     "50-gram-pamp-fortuna-gold-bar-9999-fine-sealed-in-assay000119",
}

type PureScraper struct {
}

func (p *PureScraper) GetProduct(productName string) (*Product, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var priceLabelNodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(pureURL+productMap[productName]),
		chromedp.Nodes("Highest Bid", &priceLabelNodes, chromedp.BySearch),
	)
	if err != nil {
		return nil, err
	}

	priceNode := priceLabelNodes[0].Parent.Parent.Children[1].Children[0]
	priceString := strings.Replace(priceNode.NodeValue, ",", "", -1)
	fmt.Println(priceString)

	price, err := strconv.ParseFloat(priceString[1:len(priceString)], 64)
	if err != nil {
		return nil, err
	}

	return &Product{
		productName,
		price,
	}, nil
}
