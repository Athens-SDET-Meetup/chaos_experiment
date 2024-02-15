package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Current time is: "+fmt.Sprint(time.Now().Unix()))
	})

	r.Run(":3000")
}
