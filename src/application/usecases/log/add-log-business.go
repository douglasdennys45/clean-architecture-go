package log

import (
	"github.com/douglasdennys/clean-architecture-go/src/application/ports/business"
	repo "github.com/douglasdennys/clean-architecture-go/src/application/ports/repositories/log"
	"github.com/douglasdennys/clean-architecture-go/src/domain/entities"
)

type addLogBusiness interface {
	business.AddLog
}

type app struct {
	add repo.AddLogRepository
}

func NewAddLogBusiness(add repo.AddLogRepository) addLogBusiness {
	return &app{add: add}
}

func (a *app) Add(logData []byte) error {
	result := entities.NewLog(logData)
	err := a.add.Add(result)
	return err
}
