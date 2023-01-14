package main

import (
	"fmt"
	"log"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/infrastructure/api"
	"go.uber.org/zap"
)

const DefaultPort = "8080"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	app := api.NewAPI().LoadAPI()

	logger.Info(fmt.Sprintf("Starting app on port: %s", DefaultPort))

	app.Serve(DefaultPort)
}
