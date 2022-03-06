package main

import (
	"fmt"
	"hockeyTrainer/pkg/fastHands"
	"hockeyTrainer/pkg/server"
)

func main() {
	fmt.Println("Let's Train!")
	drills, metadata := fastHands.Run()
	fmt.Println(drills)
	fmt.Println(metadata)
	server.StartServer()
}
