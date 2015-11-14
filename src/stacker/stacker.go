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

	var ho stack.Stack
	ho.Push("Xiaomi")
	ho.Push("Huawei")
	ho.Push(1999)
	x, err := ho.Top()
	fmt.Println(x)
	ho.Push(-6e-4)
	ho.Push("Apple")
	ho.Push("MS")
	x, err = ho.Top()
	fmt.Println(x)
	ho.Push(11.11)
	fmt.Println("stack is empty?", ho.IsEmpty())
	fmt.Printf("Len() == %d Cap == %d\n", ho.Len(), ho.Cap())
	diff := ho.Cap() - ho.Len()
	for i := 0; i < diff; i ++ {
		ho.Push(strings.Repeat("*", diff - i))
	}
	fmt.Printf("Len() == %d  Cap() == %d\n", ho.Len(), ho.Cap())
	for ho.Len() > 0 {
		x, _ := ho.Pop()
		fmt.Printf("%T %v\n", x, x)
	}
	fmt.Println("stack is empty?", ho.IsEmpty())
	x, err = ho.Pop()
	fmt.Println(x, err)
	x, err = ho.Pop()
	fmt.Println(x, err)
}