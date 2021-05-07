package shared

import (
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/errors"
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/ports"
)

func NoContent() ports.HttpResponse {
	return ports.HttpResponse{StatusCode: 204, Body: nil}
}

func ServerError(reason string) ports.HttpResponse {
	return ports.HttpResponse{StatusCode: 500, Body: errors.ServerError(reason)}
}