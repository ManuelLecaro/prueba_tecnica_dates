package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/ports/service"
	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func TestHolidayHandler_GetHoliday(t *testing.T) {
	type fields struct {
		name             string
		requestBody      string
		wantResponseBody string
		wantStatus       int
		service          service.HolidayService
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "should work correctly",
			fields: fields{
				service: nil,
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := "/v1/holiday"

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
