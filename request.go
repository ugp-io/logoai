package logoai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	APIKey string
	Logos  LogosService
}

func NewClient(apiKey string) *Client {

	c := &Client{
		APIKey: apiKey,
	}

	c.Logos = &LogoServiceOp{client: c}

	return c

}

func (c *Client) Request(method string, url string, bodyJSON io.Reader, response *[]byte) error {

	httpReq, errNewRequest := http.NewRequest(method, url, bodyJSON)
	if errNewRequest != nil {
		return errNewRequest
	}

	httpReq.Header.Add("LOGOAPIKEY", c.APIKey)

	client := &http.Client{}
	res, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	*response = bodyBytes

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		var errResp map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &errResp); err != nil {
			return fmt.Errorf("request failed with status %d: %s", res.StatusCode, string(bodyBytes))
		}
		return fmt.Errorf("request failed with status %d: %v", res.StatusCode, errResp)
	}

	return nil
}
