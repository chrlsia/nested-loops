package main

import (
	"fmt"
	"time"
)

var chan1=make(chan string)
var chan2=make(chan string)

func taske1(){
	time.Sleep(1*time.Second)
	chan1<-"one"
}
func taske2(){
	time.Sleep(2*time.Second)
	chan2<-"two"
}

func main(){
	go taske1()
	go taske2()

	for i:=0;i<2;i++{
		select{
		case msg1:=<-chan1:
			fmt.Println("received",msg1)
		case msg2:=<-chan2:
			fmt.Println("received",msg2)
		}
	}
}