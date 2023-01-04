package blop

import (
	"fmt"
	"log"
	"time"
)

func Run() {
	log.Println("started blop ...")
	for range time.NewTicker(3 * time.Second).C {
		fmt.Println("hello from blop", time.Now())
	}
}
