package main

import (
	"fmt"
	"stacker/stack"
	"strings"
)

func main() {
	var hi stack.Stack
	hi.Push("hi")
	hi.Push("hello")
	hi.Push(22)
	hi.Push([]string{"good", "morning", "bye"})
	hi.Push(-12.3)

	for {
		item, err := hi.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}