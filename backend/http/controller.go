package http

import (
	"bytes"
	"faker-bot/backend/entities"
	"faker-bot/backend/orm"
	"faker-bot/backend/utils"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) HotVideo() (http.Response, error) {

	access_token := utils.GetInstance().Get("access_token").(string)
	var param entities.Param
	orm.GetInstance().Model(&param).Where("id=?", 1).First(&param)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", param.Url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access-token", access_token)

	resp, err := client.Do(request)
	if err != nil {
		return http.Response{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return http.Response{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}

func (c *Controller) Topic() (http.Response, error) {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}

func (c *Controller) Amusement() (http.Response, error) {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}

func (c *Controller) Drama() (http.Response, error) {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}
func (c *Controller) Search(keyword string) (http.Response, error) {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}

func (c *Controller) HotSearch() (http.Response, error) {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}

func (c *Controller) Trending() (http.Response, error) {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}

func (c *Controller) HotSentence(hot_sentence string) (http.Response, error) {
	return http.Response{
		Status:     "success",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 0,
	}, nil
}
