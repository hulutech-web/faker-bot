package http

import "net/http"

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) HotVideo() http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}

func (c *Controller) Topic() http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}

func (c *Controller) Amusement() http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}

func (c *Controller) Drama() http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}
func (c *Controller) Search(keyword string) http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}

func (c *Controller) HotSearch() http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}

func (c *Controller) Trending() http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}

func (c *Controller) HotSentence(hot_sentence string) http.Response {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}
}
