package http

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

// 授权client_token
func (c *Client) Auth(url string, headers map[string]string) {
}
