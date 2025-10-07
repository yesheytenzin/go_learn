package main

import "fmt"

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
