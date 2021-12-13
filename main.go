package main

import "fmt"

func main(){
	fmt.Println("siannas")
	support()
	fmt.Println(saySomething())
}

func support(){
	fmt.Println("Support")
	//test()
}

func saySomething() string {
	return "Say someting"
}