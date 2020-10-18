package main

import "fmt"

func main()  {
	var todoList = [4]string{
		"123",
		"456",
		"789",
		"000",
	}

	tasks := todoList[1:4]
	for i := range tasks {
		fmt.Println(tasks[i])
	}

	todoList[3] = "asd"
	fmt.Println("\n_____________\n")
	changeTasks(tasks)

	for i := range tasks {
		fmt.Println(tasks[i])
	}
}

func changeTasks(tasks []string)  {
	tasks[0] = "test1"
	tasks[1] = "test2"
}
