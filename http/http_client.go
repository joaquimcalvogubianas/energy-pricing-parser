package http

type HttpClient interface {
	Request() HttpRequest
}

type HttpRequest interface {
	Get(url string) (RawHttpResponse, error)
}

type RawHttpResponse interface {
	Body() []byte
}
