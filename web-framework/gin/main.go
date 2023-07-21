package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello Gin")

	router := gin.New()

	router.Run(":3000")
}
