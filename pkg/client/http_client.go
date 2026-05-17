package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	httpClient = &http.Client{}

	httpUserAgent = "rtnl/blaze 0.1"
)

func HttpGetJson(url string, out any) (err error) {
	var (
		req *http.Request
		res *http.Response
	)

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", httpUserAgent)

	res, err = httpClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		err = fmt.Errorf("http response status code is %d", res.StatusCode)
		return
	}

	if res.Body == nil {
		err = fmt.Errorf("http response body is nil")
		return
	}
	defer res.Body.Close()

	reader := json.NewDecoder(res.Body)
	err = reader.Decode(out)
	if err != nil {
		return
	}

	return
}

func HttpGetDownload(url string, out io.Writer, hash *string) (err error) {
	var (
		req *http.Request
		res *http.Response
	)

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", httpUserAgent)

	res, err = httpClient.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		err = fmt.Errorf("http response status code is %d", res.StatusCode)
		return
	}

	if res.Body == nil {
		err = fmt.Errorf("http response body is nil")
		return
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return
	}

	return
}
