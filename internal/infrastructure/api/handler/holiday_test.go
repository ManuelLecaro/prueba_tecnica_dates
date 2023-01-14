package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/ports/service"
	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/ports/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestHolidayHandler_GetHoliday(t *testing.T) {
	type fields struct {
		name       string
		wantStatus int
		service    service.HolidayService
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "should work correctly",
			fields: fields{
				service: func() service.HolidayService {
					service := &mocks.HolidayService{}
					service.On("GetHolidays", mock.Anything, mock.Anything).Return([]core.Holiday{})
					return service
				}(),
				wantStatus: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := "/"

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, path, nil)
			r.Header.Add("Content-Type", gin.MIMEJSON)
			c, _ := gin.CreateTestContext(w)
			c.Request = r
			h := NewHolidayHandler(tt.fields.service)

			h.GetHoliday(c)
			assert.Equal(t, tt.fields.wantStatus, c.Writer.Status())
		})
	}
}
