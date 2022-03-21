package main

import (
	"fmt"
	"hockeyTrainer/pkg/fastHands"
	"hockeyTrainer/pkg/server"
)

func main() {
	fmt.Println("Let's Train!")
	// drills, _, _ := fastHands.Run()
	fastHands.Run()
	server.StartServer()
}
