package main

import "fmt"
/*
func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayBye(name string) {
	fmt.Println("GoodBye", name)
}

func beSocial(name string) {
	sayHello(name)
	fmt.Println("How is the weather??")
	sayBye(name)
}

func addOne(x int) int {
	return x+1
}

func sayBunch() {
	fmt.Println("Hello")
	sayBunch()
}

func main() {
	beSocial("Yeshey")
	beSocial("Tenzin")
	x := 5
	x = addOne(x)
	fmt.Println(x)


	x = addOne(addOne(x))
	fmt.Println(x)
	sayBunch()
}
*/

func sayHello(name string) {
	fmt.Println("hello",name)
}

func sayBye(name string) {
	fmt.Println("bye",name)
}

func beSocial(name string) {
	sayHello(name)
	fmt.Println("how was ur day??")
	sayBye(name)
}

func addOne(x int) int {
	return x + 1
}

// apply recursion
func sayBunch(){
	fmt.Println("hello")
	sayBunch()
}

func main() {
	beSocial("yeshey")
	sum := addOne(5)
	fmt.Println(sum)
	sayBunch()
}
