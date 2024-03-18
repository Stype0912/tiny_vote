package main

import (
	"tiny_vote/method"
	"tiny_vote/model/db"
	"tiny_vote/model/redis"
	"tiny_vote/router"
)

func main() {
	// Initialize database, redis and ticket generation.
	db.Init()
	redis.Init()
	method.Init()

	r := router.Router
	router.SetRouter()
	r.Run(":8888")
}
