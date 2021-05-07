package ports

type HttpRequest struct {
	Body []byte `json:"body"`
}
type HttpResponse struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
}

