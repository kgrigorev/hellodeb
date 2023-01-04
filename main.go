package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("started...")
	for range time.NewTicker(3 * time.Second).C {
		fmt.Println("Hello there", time.Now())
	}
}
