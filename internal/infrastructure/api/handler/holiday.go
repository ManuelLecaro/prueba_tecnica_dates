package handler

import (
	"context"
	"net/http"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/ports/service"
	"github.com/gin-gonic/gin"
)

type HolidayHandler struct {
	service service.HolidayService
}

func NewHolidayHandler(service service.HolidayService) *HolidayHandler {
	return &HolidayHandler{
		service: service,
	}
}

// GetHoliday handles the petitions for holidays data
func (hh *HolidayHandler) GetHoliday(ctx *gin.Context) {
	date, _ := ctx.GetQuery("date")
	title, _ := ctx.GetQuery("title")
	hType, _ := ctx.GetQuery("type")
	unalienable, _ := ctx.Get("inalienable")
	extra, _ := ctx.GetQuery("extra")

	unalienableValue, ok := unalienable.(bool)
	if !ok {
		unalienableValue = false
	}

	holiday := core.NewHoliday(date, title, hType, extra, unalienableValue)

	r := hh.service.GetHolidays(context.Background(), *holiday)

	ctx.JSON(http.StatusOK, r)
}
