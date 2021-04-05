package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Hammad")
	fmt.Println(message)
}
