package main

import (
	"fmt"
	"github.com/segmentio/ksuid"
)

func main() {
	id := ksuid.New().String()
	fmt.Println("I'm alive! " + id)
}
