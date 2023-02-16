package main

import (
	"fmt"

	"github.com/SaumitraLohokare/persistent_stack/pstack"
)

func main() {
	stack := pstack.NewPersistentStack[string]()
	stack.Push("A")
	stack.Push("B")
	stack.Push("C")
	stack.Push("D")
	stack.RememberPoint("point1")
	stack.Push("E")
	stack.Push("F")
	stack.Push("G")
	popped, err := stack.PeekTill("point1")
	if err != nil {
		panic(err.Error())
	}
	for i, item := range popped {
		fmt.Printf("%d: %s\n", i, item)
	}

	fmt.Println("==========")
	popped, err = stack.PopTill("point2")
	if err != nil {
		panic(err.Error())
	}
	for i, item := range popped {
		fmt.Printf("%d: %s\n", i, item)
	}

	fmt.Println("==========")
	popped = stack.PopAll()
	for i, item := range popped {
		fmt.Printf("%d: %s\n", i, item)
	}
	fmt.Println("==========")
}
