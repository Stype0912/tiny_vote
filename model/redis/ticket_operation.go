package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"log"
	"sync"
)

var rwMutex sync.RWMutex

func SetTicket(ticket string) error {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	err := rdb.Set("ticket", ticket, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetTicket() (ticket string, err error) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	ticket, err = rdb.Get("ticket").Result()
	if errors.Is(err, redis.Nil) {
		log.Println("name does not exist")
		return
	} else if err != nil {
		log.Printf("get name failed, err: %v", err)
		return
	} else {
		return
	}
}
