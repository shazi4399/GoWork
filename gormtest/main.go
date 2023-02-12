package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/gormtest/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	err := r.Run(":8082")
	if err != nil {
		log.Fatalln(err)
	}
}
