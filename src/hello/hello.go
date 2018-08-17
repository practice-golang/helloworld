package main

import (
	"fmt"
	"mypkg"

	// "spkg" // 이렇게는 안 된다.
	"vpkg"
)

func main() {
	// Hello world
	fmt.Println("Hello world!")

	var myName = "Chulsoo"
	yourName := "Younghee"

	// Using package
	fmt.Println(mypkg.SayHello())
	var name string
	name = mypkg.MyName(myName)
	fmt.Println(name)

	fmt.Println(mypkg.YourName(yourName))

	// fmt.Println(spkg.SayGoodToSeeYou()) // 이렇게는 안 된다.

	fmt.Println(vpkg.SayGoodbye())
}
