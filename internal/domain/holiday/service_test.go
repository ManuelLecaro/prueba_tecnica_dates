package holiday

import (
	"context"
	"reflect"
	"testing"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/ports/repository"
)

func TestHolidayService_GetHolidays(t *testing.T) {
	type fields struct {
		repository repository.HolidayRepository
	}
	type args struct {
		ctx     context.Context
		holiday core.Holiday
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*core.Holiday
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := &HolidayService{
				repository: tt.fields.repository,
			}
			if got := hs.GetHolidays(tt.args.ctx, tt.args.holiday); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HolidayService.GetHolidays() = %v, want %v", got, tt.want)
			}
		})
	}
}
