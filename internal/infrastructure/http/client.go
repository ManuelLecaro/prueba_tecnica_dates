package http

import (
	"context"

	"github.com/ManuelLecaro/prueba_tecnica_dates/internal/domain/core"
	"github.com/go-resty/resty/v2"
)

type Response struct {
	Status string         `json:"status"`
	Data   []core.Holiday `json:"data"`
}

type GenericClient struct {
	client *resty.Client
}

func NewGenericClient() *GenericClient {
	return &GenericClient{
		client: resty.New(),
	}
}

func (gc *GenericClient) Get(ctx context.Context, url string, body interface{}) error {
	if _, err := gc.client.R().SetResult(body).Get(url); err != nil {
		return err
	}

	return nil
}
