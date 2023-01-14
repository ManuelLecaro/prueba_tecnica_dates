package service

import (
	"context"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
)

//go:generate mockery --name HolidayService

type HolidayService interface {
	GetHolidays(ctx context.Context, holiday core.Holiday) []core.Holiday
}
