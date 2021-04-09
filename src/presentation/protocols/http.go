package protocols

type HttpResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type HttpRequest struct {
	Body        []byte      `json:"body"`
	Headers     interface{} `json:"headers"`
	Params      interface{} `json:"params"`
	QueryString interface{} `json:"queryString"`
}
