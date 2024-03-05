package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/wiggers/goexpert/desafio/temperature-ms-1/configs"
	"github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/infra/controller"
	opentel "github.com/wiggers/goexpert/desafio/temperature-ms-1/internal/infra/openTel"
)

func main() {

	config := configs.LoadConfig(".")

	url := flag.String("zipkin", config.Zipkin, "zipkin url")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	openTel := opentel.NewOpenTel()
	shutdown, err := openTel.InitTracer(*url)
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

	http.ListenAndServe(":"+config.WebServerPort, router)

}
