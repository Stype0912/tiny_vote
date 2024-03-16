package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"log"
)

func SetTicket(ticket string) error {
	err := rdb.Set("ticket", ticket, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetTicket() (ticket string, err error) {
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
