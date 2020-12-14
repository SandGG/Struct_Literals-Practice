package main

import "fmt"

type figure struct {
	height, width int
	name          string
}

var (
	rectangle = figure{5, 10, "Rectangle"}
	figure1   = figure{height: 15}
	figure2   = figure{name: "Pending"}
	figure3   = figure{name: "Name", width: 20, height: 10}
)

func main() {
	fmt.Println(rectangle, figure1, figure2, figure3)
	area(rectangle)
	area(figure3)
}

func area(shape figure) {
	area := shape.height * shape.width
	fmt.Println(shape.name, area, "area")
}
