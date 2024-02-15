package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/consuming", func(ctx *gin.Context) {

		response, err := http.Get(fmt.Sprintf("http://%s:%s/", os.Getenv("PROVIDER_URL"), os.Getenv("PROVIDER_PORT")))
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}

		if response.StatusCode == 200 {
			ctx.String(200, "Successfully consuming message")
		} else {
			ctx.String(500, "Internal server error")
		}

	})

	r.Run(":3000")
}
