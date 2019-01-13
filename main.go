package main

import (
	"fmt"
	"github.com/sislow/angryMonkey/routes"
	"time"
)

var timer = time.Now()

func main() {
	fmt.Println("Boot time: ", timer)
	// or route
	routes.Router()
}
