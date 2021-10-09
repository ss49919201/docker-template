package main

import (
	"fmt"
	"time"
)

const (
	defaultMessage = "こんにちは、世界"
)

func main() {
	now := time.Now()

	fmt.Println(now.String())
	fmt.Println(defaultMessage)
}
