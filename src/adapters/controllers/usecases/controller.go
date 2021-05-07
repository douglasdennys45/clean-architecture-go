package usecases

import "github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/ports"

type Controller interface {
	Handle(req ports.HttpRequest) ports.HttpResponse
}