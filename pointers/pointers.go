package main

import "fmt"

type position struct {
	x float32
	y float32
}

type enemy struct {
	name   string
	health int32
	pos    position
}

func whereIsEnemy(b *enemy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func addOne(num *int) {
	*num = *num + 1
}

func main() {
	/*
		x := 5

		// var xPtr *int = &x // Gets the memory address location for the variable
		xPtr := &x // Shorthand

		fmt.Println(xPtr)

		addOne(xPtr)
		fmt.Println(x)
		//*/

	p := position{35, 26}

	b := enemy{"Frank", 100, p}

	whereIsEnemy(&b)
}
