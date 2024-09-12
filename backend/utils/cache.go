package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	once sync.Once
)
var cacheIns = &Cache{}

type Cache struct {
	// TODO: implement cache
	cache map[string]interface{}
}

func GetInstance() *Cache {
	return cacheIns
}

func Boot() *Cache {
	once.Do(func() {
		cacheIns = &Cache{
			cache: make(map[string]interface{}),
		}
	})
	return cacheIns
}

func (c *Cache) Get(key string) interface{} {
	return c.cache[key]
}

func (c *Cache) Set(key string, value interface{}, expire time.Duration) {
	//过期之后删除
	go func() {
		time.Sleep(expire)
		delete(c.cache, key)
	}()
	c.cache[key] = value
}

type NormalReq struct {
	Data    DataInfo `json:"data"`
	Message string   `json:"message"`
}

type DataInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Captcha     string `json:"captcha"`
	DescUrl     string `json:"desc_url"`
	Description string `json:"description"`
	ErrorCode   int    `json:"error_code"`
	LogId       string `json:"log_id"`
}

func (c *Cache) Authorize() error {
	//判断有没有数据，如果没有数据从新请求，如果有数据，直接返回
	if c.Get("access_token") != nil {
		return nil
	}
	remote_url := "https://open.douyin.com/oauth/client_token"
	//encodedKeyword := url.QueryEscape(keyword)
	//	http.NewRequest发送一个Get请求
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	body := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(body).Encode(map[string]string{
		"client_key":    "aw3xp2y63c5ezjmw",
		"client_secret": "281106cf7912a04cebf0690ea2a3180d",
		"grant_type":    "client_credential",
	})
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", remote_url, body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	//如果状态码不为200，打印状态码并退出
	if resp.StatusCode != 200 {
		return fmt.Errorf("请求失败，状态码为：%v", resp.StatusCode)
	}
	//关闭请求
	defer resp.Body.Close()
	resultBuffer := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBuffer, resp.Body)
	fmt.Println(resultBuffer.String())
	var result NormalReq
	err = json.Unmarshal(resultBuffer.Bytes(), &result)
	if err != nil {
		return err
	}
	// 2 个小时过期

	cacheIns.Set("access_token", result.Data.AccessToken, 2*time.Hour)
	return nil
}
