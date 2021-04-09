package protocols

type Middleware interface {
	Handle(httpRequest HttpRequest) HttpResponse
}