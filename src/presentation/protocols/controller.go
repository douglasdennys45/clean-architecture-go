package protocols

type Controller interface {
	Handle(httpRequest HttpRequest) HttpResponse
}