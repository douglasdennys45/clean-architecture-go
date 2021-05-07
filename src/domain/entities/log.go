package entities

import (
	"encoding/json"
	"github.com/douglasdennys/clean-architecture-go/src/domain/data"
	"time"

	"github.com/go-playground/validator/v10"
)

type Log struct {
	Error     string    `json:"name" validate:"required"`
	Service   string    `json:"service" validate:"required"`
	Uri       string    `json:"uri" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

func NewLog(logData []byte) []byte {
	var log data.LogData
	_ = json.Unmarshal(logData, &log)
	entity := Log{log.Error, log.Service, log.Uri, time.Now()}
	_ = entity.isValid()
	result, _ := json.Marshal(entity)
	return result
}

func (l *Log) isValid() error {
	validate := validator.New()
	err := validate.Struct(l)
	if err != nil {
		return err
	}
	return nil
}
