package wandgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	BaseURL = "http://melpon.org/wandbox/api"
)

type API struct {
	HTTPClient *http.Client
}

// NewAPI returns a API struct
func NewAPI() *API {
	api := &API{
		HTTPClient: http.DefaultClient,
	}

	return api
}

func (a *API) apiGet(url string, v url.Values, data interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = v.Encode()

	resp, err := a.HTTPClient.Do(req)
	defer resp.Body.Close()

	return decode(resp, data)
}

func (a *API) apiPost(url string, v url.Values, data interface{}) error {
	resp, err := a.HTTPClient.PostForm(url, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decode(resp, data)
}

func (a *API) apiPostAsJSON(url string, v interface{}, data interface{}) error {
	j, err := json.Marshal(v)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decode(resp, data)
}

func decode(resp *http.Response, data interface{}) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %s", resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
