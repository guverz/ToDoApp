package main

func main() {
	todos := Todos{}
	todos.add("Cook the chicken")
	todos.add("Mop the floor")
	todos.add("Clean the dishes")

	todos.del(1)
	todos.toggle(1)
	todos.edit(0, "Cook the turkey")
	todos.print()
}
