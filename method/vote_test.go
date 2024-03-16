package method

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
	"tiny_vote/model/redis"
)

func TestVoteWithinValidTime(t *testing.T) {
	t.Parallel()
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	redis.Init()

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
	q.Add("query", fmt.Sprintf(`mutation{vote(name:["Alice", "Bob"], ticket:"%v")}`, ticket))
	req.URL.RawQuery = q.Encode()

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	respBody, _ = io.ReadAll(resp.Body)
	t.Logf("Vote result: %v", string(respBody))
	respJson, err = simplejson.NewJson(respBody)
	if err != nil {
		t.Error(err)
	}
	voteResult, _ := respJson.Get("data").Get("vote").Bool()
	assert.Equal(t, true, voteResult)
}

func TestVoteOutOfValidTime(t *testing.T) {
	t.Parallel()
	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	redis.Init()

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

	time.Sleep(2 * time.Second)

	req, err = http.NewRequest("GET", "http://localhost:8888/graphql", nil)
	q = req.URL.Query()
	q.Add("query", fmt.Sprintf(`mutation{vote(name:["Alice", "Bob"], ticket:"%v")}`, ticket))
	req.URL.RawQuery = q.Encode()

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	respBody, _ = io.ReadAll(resp.Body)
	t.Logf("Vote result: %v", string(respBody))
	respJson, err = simplejson.NewJson(respBody)
	if err != nil {
		t.Error(err)
	}
	voteResult, _ := respJson.Get("data").Get("vote").Bool()
	voteMessage, _ := respJson.Get("errors").GetIndex(0).Get("message").String()
	assert.Equal(t, false, voteResult)
	assert.Equal(t, "your ticket has been expired", voteMessage)
}
