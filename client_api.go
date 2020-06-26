package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// BaseURL ...
const BaseURL = `https://api.chatwork.com/v2`

// ClientAPI ...
type ClientAPI struct {
	APIToken string
	BaseURL  string
}

// NewCientAPI ...
func NewCientAPI(apiToken string) *ClientAPI {
	return &ClientAPI{APIToken: apiToken, BaseURL: BaseURL}
}

// Get ...
func (c *ClientAPI) Get(endpoint string, params map[string]string) []byte {
	return c.sendRequest("GET", endpoint, params)
}

// Post ...
func (c *ClientAPI) Post(endpoint string, params map[string]string) []byte {
	return c.sendRequest("POST", endpoint, params)
}

// Put ...
func (c *ClientAPI) Put(endpoint string, params map[string]string) []byte {
	return c.sendRequest("PUT", endpoint, params)
}

// Delete ...
func (c *ClientAPI) Delete(endpoint string, params map[string]string) []byte {
	return c.sendRequest("DELETE", endpoint, params)
}

func (c *ClientAPI) buildURL(baseURL, endpoint string, params map[string]string) string {
	query := make([]string, len(params))
	for k := range params {
		query = append(query, k+"="+params[k])
	}
	return baseURL + endpoint + "?" + strings.Join(query, "&")
}

func (c *ClientAPI) buildBody(params map[string]string) url.Values {
	body := url.Values{}
	for k := range params {
		body.Add(k, params[k])
	}
	return body
}

func (c *ClientAPI) parseBody(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte(``)
	}
	return body
}

func (c *ClientAPI) sendRequest(method, endpoint string, params map[string]string) []byte {
	httpClient := &http.Client{}

	var (
		req        *http.Request
		requestErr error
	)

	if method != "GET" {
		req, requestErr = http.NewRequest(method, c.BaseURL+endpoint, bytes.NewBufferString(c.buildBody(params).Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, requestErr = http.NewRequest(method, c.buildURL(c.BaseURL, endpoint, params), nil)
	}
	if requestErr != nil {
		panic(requestErr)
	}

	req.Header.Add("X-ChatWorkToken", c.APIToken)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		return []byte(``)
	}

	return c.parseBody(resp)
}
