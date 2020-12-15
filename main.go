package main

import "fmt"

type figure struct {
	height, width int
	name          string
}

func main() {
	area(figure{name: "Name", width: 20, height: 10})
	area(figure{name: "Pending", width: 100})
}

func area(shape figure) {
	area := shape.height * shape.width
	fmt.Println(shape.name, area, "area")
}
