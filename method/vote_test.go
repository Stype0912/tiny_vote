package method

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tiny_vote/method/http_request"
	"tiny_vote/model/redis"
)

func TestVoteWithinValidTime(t *testing.T) {
	t.Parallel()
	var err error
	var respBody []byte

	redis.Init()

	respBody, err = http_request.GraphqlRequest("http://localhost:8888/graphql", "query{cas}")
	if err != nil {
		t.Error(err)
	}
	respJson, err := simplejson.NewJson(respBody)
	if err != nil {
		t.Error(err)
	}
	ticket, _ := respJson.Get("data").Get("cas").String()
	respBody, err = http_request.GraphqlRequest("http://localhost:8888/graphql", fmt.Sprintf(`mutation{vote(name:["Alice", "Bob"], ticket:"%v")}`, ticket))
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
	var respBody []byte

	redis.Init()

	respBody, err = http_request.GraphqlRequest("http://localhost:8888/graphql", "query{cas}")
	if err != nil {
		t.Error(err)
	}
	respJson, err := simplejson.NewJson(respBody)
	if err != nil {
		t.Error(err)
	}
	ticket, _ := respJson.Get("data").Get("cas").String()

	time.Sleep(refreshInterval)

	respBody, err = http_request.GraphqlRequest("http://localhost:8888/graphql", fmt.Sprintf(`mutation{vote(name:["Alice", "Bob"], ticket:"%v")}`, ticket))
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
