package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type application struct {
	config  config
	scraper Scraper
	http.Handler
}

type config struct {
	port int
	cors struct {
		trustedOrigins []string
	}
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()

	app := &application{
		config:  cfg,
		scraper: new(PureScraper),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	fmt.Printf("listening on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}

}
