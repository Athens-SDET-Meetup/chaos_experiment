package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/consuming", func(ctx *gin.Context) {
		retryTimes := 0

		for i := 1; i < 5; i++ {

			response, err := http.Get(fmt.Sprintf("http://%s:%s/", os.Getenv("PROVIDER_URL"), os.Getenv("PROVIDER_PORT")))
			if err != nil {
				fmt.Printf("error making http request: %s\n", err)
				os.Exit(1)
			}

			if response.StatusCode == 200 {
				break
			} else if response.StatusCode == 500 {
				retryTimes += 1
			}

			time.Sleep(2 * time.Second)

		}

		if retryTimes >= 4 {
			ctx.String(500, "Internal server error")
		} else {
			ctx.String(200, "Successfully consuming message")
		}

	})

	r.Run(":3000")
}
