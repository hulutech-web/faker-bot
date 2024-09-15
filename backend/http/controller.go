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

func (c *Controller) HotVideo() (bytes.Buffer, error) {

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
		return bytes.Buffer{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return bytes.Buffer{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	return *resultBuffer, nil
}

func (c *Controller) Topic() (bytes.Buffer, error) {
	access_token := utils.GetInstance().Get("access_token").(string)
	var param entities.Param
	orm.GetInstance().Model(&param).Where("id=?", 2).First(&param)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", param.Url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access-token", access_token)

	resp, err := client.Do(request)
	if err != nil {
		return bytes.Buffer{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return bytes.Buffer{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	return *resultBuffer, nil
}

func (c *Controller) Amusement() (bytes.Buffer, error) {
	access_token := utils.GetInstance().Get("access_token").(string)
	var param entities.Param
	orm.GetInstance().Model(&param).Where("id=?", 3).First(&param)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", param.Url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access-token", access_token)

	resp, err := client.Do(request)
	if err != nil {
		return bytes.Buffer{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return bytes.Buffer{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	return *resultBuffer, nil
}

func (c *Controller) Drama() (bytes.Buffer, error) {
	access_token := utils.GetInstance().Get("access_token").(string)
	var param entities.Param
	orm.GetInstance().Model(&param).Where("id=?", 4).First(&param)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", param.Url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access-token", access_token)

	resp, err := client.Do(request)
	if err != nil {
		return bytes.Buffer{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return bytes.Buffer{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	return *resultBuffer, nil
}

func (c *Controller) HotSearch() (bytes.Buffer, error) {
	access_token := utils.GetInstance().Get("access_token").(string)
	var param entities.Param
	orm.GetInstance().Model(&param).Where("id=?", 5).First(&param)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", param.Url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access-token", access_token)

	resp, err := client.Do(request)
	if err != nil {
		return bytes.Buffer{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return bytes.Buffer{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	return *resultBuffer, nil
}

func (c *Controller) Trending() (bytes.Buffer, error) {
	access_token := utils.GetInstance().Get("access_token").(string)
	var param entities.Param
	orm.GetInstance().Model(&param).Where("id=?", 6).First(&param)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", param.Url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access-token", access_token)

	resp, err := client.Do(request)
	if err != nil {
		return bytes.Buffer{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return bytes.Buffer{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	return *resultBuffer, nil
}

func (c *Controller) HotSentence(hot_sentence string) (bytes.Buffer, error) {
	access_token := utils.GetInstance().Get("access_token").(string)
	var param entities.Param
	orm.GetInstance().Model(&param).Where("id=?", 7).First(&param)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", param.Url, nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("access-token", access_token)

	resp, err := client.Do(request)
	if err != nil {
		return bytes.Buffer{}, err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return bytes.Buffer{}, fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	return *resultBuffer, nil
}
