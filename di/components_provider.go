package di

import "github.com/go-resty/resty/v2"

func GetHttpClient() *resty.Client {
	return resty.New()
}
