package http

import "github.com/douglasdennys/go-mongodb/src/presentation/protocols"

func BadRequest(err error) protocols.HttpResponse {
	return protocols.HttpResponse{400, err}
}

func Forbidden(err error) protocols.HttpResponse  {
	return protocols.HttpResponse{403, err}
}

func Unauthorized() protocols.HttpResponse {
	return protocols.HttpResponse{401, "Unauthorized"}
}

func ServerError(err error) protocols.HttpResponse {
	return protocols.HttpResponse{500, err}
}

func Ok(data interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{200, data}
}

func Created(data interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{201, data}
}

func NoContent() protocols.HttpResponse {
	return protocols.HttpResponse{204, ""}
}
