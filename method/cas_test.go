package method

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
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

			time.Sleep(2 * time.Second)

			resp, err = http.DefaultClient.Do(req)
			if err != nil {
				t.Error(err)
			}
			respBody, _ = io.ReadAll(resp.Body)
			secondTicket := string(respBody)
			t.Logf("Second Ticket: %v", secondTicket)

			assert.NotEqual(t, firstTicket, secondTicket)
		}()
	}
	wg.Wait()
}
