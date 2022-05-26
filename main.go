package main

import (
	"assignment-2/routers"
	"fmt"
)

func main() {
	fmt.Println("starting server")
	routers.Router().Run(":3000")
}
