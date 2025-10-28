package main

import "fmt"

/*
func addOne(num int) {
	num = num + 1
}

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name	string
	health	int
	pos		position
}

func whereIsBad(b *badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(",x,":",y,")")
}

func addOne(num *int) {
	// dereference a pointer
	*num = *num + 1
}

func main(){
	x := 5
	fmt.Println(x)

	//xPtr := &x
	var xPtr *int = &x
	fmt.Println(xPtr)

	// when you pass a variable to a function its just a copy. function add but does
	// not do anything to original value instead use pointer
	addOne(xPtr)
	fmt.Println(x)

	p := position{3, 2}
	b := badGuy{"yeshey",1000,p}
	whereIsBad(&b)
}

*/
func addOne(num *int){
	*num = *num + 1
}

func main() {
	x := 5
	xPtr := &x
	fmt.Println(xPtr)
	addOne(xPtr)
	fmt.Println(x)
}



