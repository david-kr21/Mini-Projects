package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("localTransfer started . . .")
	fmt.Println("What do you want to transfer?")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"message": input,
		})
	})

	router.POST("/ack", func(ctx *gin.Context) {
		os.Exit(0)
	})

	router.Run("0.0.0.0:4956")
}
