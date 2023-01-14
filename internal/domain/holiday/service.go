package holiday

import (
	"context"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/ports/repository"
)

type HolidayService struct {
	repository repository.HolidayRepository
}

func NewHolidayService(repo repository.HolidayRepository) *HolidayService {
	return &HolidayService{
		repository: repo,
	}
}

func (hs *HolidayService) GetHolidays(ctx context.Context, holiday core.Holiday) []core.Holiday {
	return hs.repository.GetHoliday(ctx, holiday)
}
