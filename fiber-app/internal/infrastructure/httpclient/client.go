package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL string
}

func New(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

func (c *Client) Post(endpoint string, body interface{}, result interface{}) error {
	jsonData, _ := json.Marshal(body)
	resp, err := http.Post(c.baseURL+endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}

// PostWithHeaders permite enviar una petición POST con headers personalizados
func (c *Client) PostWithHeaders(endpoint string, body interface{}, headers map[string]string, result interface{}) error {
	jsonData, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", c.baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("Response status:", resp)
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}
