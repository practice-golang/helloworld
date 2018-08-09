package main

import (
	"fmt"
	"mypkg"
)

func main() {
	// Hello world
	fmt.Println("Hello world!")
	fmt.Println(mypkg.SayHello())

	// Using package
	var myName = "Chulsoo"
	yourName := "Younghee"

	var name string
	name = mypkg.MyName(myName)
	fmt.Println(name)

	fmt.Println(mypkg.YourName(yourName))
}
