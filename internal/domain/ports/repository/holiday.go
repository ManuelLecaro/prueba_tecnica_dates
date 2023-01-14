package repository

import (
	"context"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
)

//go:generate mockery --name HolidayRepository

type HolidayRepository interface {
	GetHoliday(ctx context.Context, holiday core.Holiday) []core.Holiday
}
