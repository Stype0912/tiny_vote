package main

import (
	_ "tiny_vote/method"
	"tiny_vote/model/db"
	"tiny_vote/model/redis"
	"tiny_vote/router"
)

func main() {
	db.Init()
	redis.Init()
	r := router.Router
	router.SetRouter()
	r.Run(":8888")
}
