package main

import "fmt"
/*
type position struct {
	x float32
	y float32
}

type badGuy struct {
	name	string
	health	int
	pos		position
}

func whereIsBad(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(",x,":",y,")")
}

func main() {
	//var p Position
	p := position{3, 2}
	fmt.Println(p.x)

	b := badGuy{"yeshey",1000,p}
	fmt.Println(b)
	whereIsBad(b)
}
*/

type position struct {
	x float32
	y float32
}

type badguy struct {
	name string
	health int
	pos position
}

func whereisBad(b badguy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(",x,y,")")
}

func main() {
	p := position{3,4}
	b := badguy{"yeshey",100,p}
	fmt.Println(p.y)
	fmt.Println(b.health)
	whereisBad(b)
}

