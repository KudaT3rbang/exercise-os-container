package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting infinite time loop")
	for {
		now := time.Now()
		fmt.Printf("Current Time: %s\n", now.Format("15:04:05"))
		time.Sleep(1 * time.Second)
	}
}
