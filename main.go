package main

import (
	"github.com/codecodify/chat/router"
	"log"
)

func main() {
	r := router.Router()
	log.Fatal(r.Run(":8000"))
}
