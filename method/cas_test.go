package method

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
	"tiny_vote/method/http_request"
	"tiny_vote/model/redis"
)

func TestCas(t *testing.T) {
	t.Parallel()
	redis.Init()
	var testCounts int = 100
	var wg sync.WaitGroup
	wg.Add(testCounts)
	for i := 0; i < testCounts; i++ {
		go func() {
			defer wg.Done()
			var err error
			var respBody []byte
			respBody, err = http_request.GraphqlRequest("http://localhost:8888/graphql", "query{cas}")
			if err != nil {
				t.Error(err)
			}
			firstTicket := string(respBody)
			t.Logf("First Ticket: %v", firstTicket)

			time.Sleep(refreshInterval)

			respBody, err = http_request.GraphqlRequest("http://localhost:8888/graphql", "query{cas}")
			if err != nil {
				t.Error(err)
			}
			secondTicket := string(respBody)
			t.Logf("Second Ticket: %v", secondTicket)

			assert.NotEqual(t, firstTicket, secondTicket)
		}()
	}
	wg.Wait()
}
