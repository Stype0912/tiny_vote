package http_request

import (
	"io"
	"net/http"
)

func GraphqlRequest(url string, body string) (respBody []byte, err error) {
	var req *http.Request
	var resp *http.Response

	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	q := req.URL.Query()
	q.Add("query", body)
	req.URL.RawQuery = q.Encode()

	resp, err = http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	respBody, _ = io.ReadAll(resp.Body)
	return
}
