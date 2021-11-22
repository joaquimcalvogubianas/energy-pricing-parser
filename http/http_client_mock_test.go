package http

type MockHttpCient struct {
	Req HttpRequest
}

func (mhttp MockHttpCient) Request() HttpRequest {
	return mhttp.Req
}

type MockHttpRequest struct {
	Response RawHttpResponse
}

func (httpR MockHttpRequest) Get(url string) (RawHttpResponse, error) {
	return httpR.Response, nil
}

type MockRawHttpResponse struct {
	RawResponse []byte
}

func (rawResp MockRawHttpResponse) Body() []byte {
	return rawResp.RawResponse
}
