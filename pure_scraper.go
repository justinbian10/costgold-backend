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

type PureScraper struct {
}

func (p *PureScraper) GetProduct(productName string) (*Product, error) {
	//ctx, cancel := chromedp.NewContext(context.Background())
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true), // VERY DANGEROUS!
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
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
