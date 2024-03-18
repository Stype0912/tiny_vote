package method

import (
	"log"
	"math/rand"
	"time"
	"tiny_vote/model/redis"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var refreshInterval = 2 * time.Second

func Init() {
	go generateTicket()
}

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateTicket() {
	for {
		time.Sleep(refreshInterval)
		err := redis.SetTicket(randStr(20))
		if err != nil {
			log.Println(err)
		}
	}
}
