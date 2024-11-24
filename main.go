package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Cook the chicken")
	todos.add("Mop the floor")
	todos.add("Clean the dishes")
	//
	fmt.Printf("%+v\n\n", todos)
	todos.del(1)
	fmt.Printf("%+v\n\n", todos)
}
