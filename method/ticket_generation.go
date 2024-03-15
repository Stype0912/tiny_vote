package method

import (
	"math/rand"
	"sync"
	"time"
)

type Ticket struct {
	mutex      sync.Mutex
	ticketId   string
	expiration time.Time
}

var CurrentTicket Ticket

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	go GenerateTicket()
}

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateTicket() {
	for {
		time.Sleep(2 * time.Second)
		CurrentTicket.mutex.Lock()
		CurrentTicket.ticketId = randStr(20)
		CurrentTicket.expiration = time.Now().Add(2 * time.Second)
		CurrentTicket.mutex.Unlock()
	}
}
