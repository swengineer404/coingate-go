package coingate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type apiClient struct {
	client *http.Client
	url    string
	key    string
}

type apiResponse struct {
}

type apiError struct {
	Message string   `json:"message"`
	Reason  string   `json:"reason"`
	Errors  []string `json:"errors"`
}

func (a *apiError) Error() string {
	return a.Message
}

func newAPIClient(url string, key string) *apiClient {
	return &apiClient{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		url: url,
		key: key,
	}
}

func (c *apiClient) send(path, method string, params, out any) error {
	b, _ := json.Marshal(params)
	req, err := http.NewRequest(method, c.url+path, bytes.NewReader(b))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "coingate-go")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.key))

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var resErr apiError
		if err := json.NewDecoder(res.Body).Decode(&resErr); err != nil {
			return err
		}

		return &resErr
	}

	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return err
	}

	return nil
}
