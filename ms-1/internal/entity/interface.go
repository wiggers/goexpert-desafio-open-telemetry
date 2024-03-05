package entity

import "context"

type temperatureMs2 struct {
	City string `json:"city"`
}

type AdapterInterface interface {
	FindData(ctx context.Context, zipcode *ZipCode) (*Temperature, error)
}
