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

func whereIsEnemy(b enemy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {
	p := position{35, 26}

	b := enemy{"Frank", 100, p}

	fmt.Println(b)
	whereIsEnemy(b)
}
