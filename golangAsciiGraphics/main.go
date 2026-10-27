package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var debugMode = true

var player = "@"
var playerPos = map[string]int{"height": 15, "width": 15}
var wall = "#"
var width = 50
var height = 30

var inGame = true

var gameGrid = make([][]string, height)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Platform unsupported!")
	}
}

func init() {
	for i := range gameGrid {
		gameGrid[i] = make([]string, width)
	}
}

func RenderImage() {

	CallClear()

	// Build grid
	fmt.Println("Controls:\n Left : Move player left \n Right : Move player right \n Up : Move player up \n Down : Move player down \n Space : Stop the Game \n\n")

	// Debug infos
	if debugMode {
		fmt.Println("[DEBUG INFOS]")
		fmt.Println("Player Width: " + strconv.Itoa(playerPos["width"]))
		fmt.Println("Player height: " + strconv.Itoa(playerPos["height"]))
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || i == height-1 || j == 0 || j == width-1 {
				gameGrid[i][j] = wall
			} else if i == playerPos["height"] && j == playerPos["width"] {
				gameGrid[i][j] = player
			} else {
				gameGrid[i][j] = " "
			}
		}
	}

	// render actual image
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(gameGrid[i][j])
		}
		fmt.Println()
	}
}

func main() {

	fmt.Println("[*] Starting Game Engine...")
	fmt.Println("[*] Rendering first image . . .")
	RenderImage()

	for inGame {
		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			if key.Code == keys.Right && playerPos["width"] != width-2 {
				playerPos["width"] += 1
				return true, nil
			} else if key.Code == keys.Left && playerPos["width"] != 1 {
				playerPos["width"] -= 1
				return true, nil
			} else if key.Code == keys.Up && playerPos["height"] != 1 {
				playerPos["height"] -= 1
				return true, nil
			} else if key.Code == keys.Down && playerPos["height"] != height-2 {
				playerPos["height"] += 1
				return true, nil
			} else if key.Code == keys.Space {
				inGame = false
				return true, nil
			} else {
				return false, nil
			}
		})

		RenderImage()
	}

}
