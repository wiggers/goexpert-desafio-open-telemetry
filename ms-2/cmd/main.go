package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/wiggers/goexpert/desafio/1-temperatura/configs"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/controller"
	opentel "github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/openTel"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
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
	router.Use(otelmux.Middleware(openTel.ServiceName))
	router.HandleFunc("/temperature", controller.FindTemperature)

	http.ListenAndServe(":"+config.WebServerPort, router)

}
