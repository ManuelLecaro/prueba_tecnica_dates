package repo

import (
	"context"
	"fmt"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
)

type HolidayRepository struct {
	Holidays []core.Holiday
}

func NewHolidayRepository(data []core.Holiday) *HolidayRepository {
	return &HolidayRepository{
		Holidays: data,
	}
}

func (hr *HolidayRepository) GetHoliday(ctx context.Context, h core.Holiday) []core.Holiday {
	relevants := []core.Holiday{}

	//fmt.Println("DATA::: ", hr.Holidays)

	for _, day := range hr.Holidays {
		if h.Compare(day) {
			fmt.Println("DAY:: ", day)
			relevants = append(relevants, day)
		}
	}

	return relevants
}
