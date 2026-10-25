package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func showExposedIps(port string) {
	var ipv4Pattern = regexp.MustCompile(`\b(?:(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])/(?:3[0-2]|[12]?[0-9])\b`)
	interfaceAddr, intrErr := net.InterfaceAddrs()
	if intrErr != nil {
		fmt.Println("[*] Failed to fetch interface addresses!")
	} else {
		fmt.Println("\n\nThe Transfer Server is available from the following addresses: ")
		for i := 0; i < len(interfaceAddr); i++ {
			var ipv4Str = net.Addr.String(interfaceAddr[i])
			if ipv4Pattern.MatchString(ipv4Str) {
				ipv4Str, _, _ := strings.Cut(net.Addr.String(interfaceAddr[i]), "/")
				fmt.Println("[*] " + ipv4Str + ":" + port)
			}
		}
	}
}

func main() {

	var port = "4956"

	fmt.Println("[*] localTransfer started . . .")

	if len(os.Args[1:]) <= 0 {
		fmt.Println("[*] No Arguments given, switching to manual mode . . .")
		fmt.Println("[X] What do you want to transfer?")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[*] An error occured while reading input. Please try again", err)
			return
		}

		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.LoadHTMLGlob("templates/*")

		router.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", gin.H{
				"message": input,
			})
		})

		router.POST("/ack", func(ctx *gin.Context) {
			os.Exit(0)
		})

		showExposedIps(port)

		router.Run("0.0.0.0:" + port)
	} else {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		router.LoadHTMLGlob("templates/*")

		router.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", gin.H{
				"message": strings.Join(os.Args[1:], " "),
			})
		})

		router.POST("/ack", func(ctx *gin.Context) {
			os.Exit(0)
		})

		showExposedIps(port)

		router.Run("0.0.0.0:" + port)
	}

}
