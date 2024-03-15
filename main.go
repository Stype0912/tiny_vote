package main

import (
	"tiny_vote/model"
	"tiny_vote/router"
)

func main() {
	model.Init()
	r := router.Router
	router.SetRouter()
	r.Run(":8888")
}
