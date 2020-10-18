package main

import "fmt"

func main()  {
	var todoList  = [...]string{
		"one",
		"two",
		"three",
	}

	fmt.Printf("Кол элем в списке: %d\n", len(todoList))

	for _, item := range todoList {
		fmt.Printf("%s\n", item)
	}
}


