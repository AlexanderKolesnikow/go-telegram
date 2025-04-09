package telegram

import "fmt"

type Client struct {
	Token  string
	apiURL string
}

func NewClient(token string) *Client {
	return &Client{
		Token:  token,
		apiURL: fmt.Sprintf("https://api.telegram.org/bot%s", token),
	}
}
