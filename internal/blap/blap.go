package blap

import (
	"fmt"
	"log"
	"time"
)

func Run() {
	log.Println("started blap ...")
	for range time.NewTicker(3 * time.Second).C {
		fmt.Println("hello from blap", time.Now())
	}
}
