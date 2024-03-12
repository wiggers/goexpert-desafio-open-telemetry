package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wiggers/goexpert/desafio/1-temperatura/configs"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/controller"
	opentel "github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/openTel"
)

func main() {

	config := configs.LoadConfig(".")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	openTel := opentel.NewOpenTel()

	shutdown, err := openTel.InitProvider("temperature-ms-2", config.UrlCollectior)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	controller := controller.NewFindTemperatureByZipCodeController()

	router := mux.NewRouter()
	router.HandleFunc("/temperature", controller.FindTemperature)
	router.Handle("/metrics", promhttp.Handler())

	go http.ListenAndServe(":"+config.WebServerPort, router)

	select {
	case <-sigCh:
		log.Println("ctrl+C pressed..")
	case <-ctx.Done():
		log.Println("Shutting down")
	}

}
