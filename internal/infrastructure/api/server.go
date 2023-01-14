package api

import (
	"context"
	"fmt"
	"log"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/holiday"
	hh "github.com/ManuelLecaro/prueba_tecnica_dates/internal/infrastructure/api/handler"
	genericHTTP "github.com/ManuelLecaro/prueba_tecnica_dates/internal/infrastructure/http"
	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/infrastructure/repo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	APIName = "holidays_app"
	urlData = "https://api.victorsanmartin.com/feriados/en.json"
)

type Api struct {
	router     *gin.Engine
	handler    *hh.HolidayHandler
	logger     *zap.Logger
	httpClient *genericHTTP.GenericClient
}

func NewAPI() *Api {
	gin.SetMode(gin.ReleaseMode)

	api := &Api{
		router: gin.New(),
	}

	api.router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{
				fmt.Sprintf("/api/%s/health-check", APIName),
			},
		}),
		gin.Recovery(),
	)

	return api
}

func (api *Api) LoadAPI() *Api {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	api.logger = logger

	ctx := context.Background()

	client := genericHTTP.NewGenericClient()
	api.httpClient = client

	api.logger.Info("Loading Info")

	baseData := &genericHTTP.Response{}
	if err := api.httpClient.Get(ctx, urlData, baseData); err != nil {
		api.logger.Error(err.Error())
	}

	api.logger.Info("Info Loaded successfully")

	fmt.Println("DATA::: ", baseData.Data)

	repo := repo.NewHolidayRepository(baseData.Data)

	service := holiday.NewHolidayService(repo)
	handler := hh.NewHolidayHandler(service)
	api.handler = handler

	api.startRouting()

	return api
}

func (api *Api) Serve(port string) {
	api.router.Run(fmt.Sprintf(":%s", port))
}
