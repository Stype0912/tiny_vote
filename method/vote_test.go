package method

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestVote(t *testing.T) {
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	req, err = http.NewRequest("GET", "http://localhost:8888/graphql", nil)
	if err != nil {
		t.Error(err)
	}
	q := req.URL.Query()
	q.Add("query", "query{cas}")
	req.URL.RawQuery = q.Encode()

	resp, err = http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		t.Error(err)
	}
	respBody, _ = io.ReadAll(resp.Body)
	firstTicket := string(respBody)
	t.Logf("First Ticket: %v", firstTicket)

	respJson, err := simplejson.NewJson(respBody)
	if err != nil {
		t.Error(err)
	}
	ticket, _ := respJson.Get("data").Get("cas").String()

	req, err = http.NewRequest("GET", "http://localhost:8888/graphql", nil)
	q = req.URL.Query()
	log.Println(fmt.Sprintf(`vote(name:["Alice", "Bob"], ticket:"%v")`, ticket))
	q.Add("query", fmt.Sprintf(`mutation{vote(name:["Alice", "Bob"], ticket:"%v")}`, ticket))
	req.URL.RawQuery = q.Encode()

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	respBody, _ = io.ReadAll(resp.Body)
	t.Logf("Vote result: %v", string(respBody))
}
