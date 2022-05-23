package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader=bufio.NewReadWriter(os.Stdin)

func main(){
	firstName:= getUserData("Please enter your first name:")
	lastName:= getUserData("Please enter your last name:")
	birthdate:= getUserData("Please enter your birthdate (MM/DD/YY):")
	created := time.Now()

	// ... do something awesome with these data
	fmt.Println(firstName, lastName,birthdate,created)
}

func getUserData(promptText string) string{
	fmt.Print(promptText)

	userInput, _:= reader.ReadString('\n')
	cleanedInput:= strings.Replace(userInput,"\n","",-1)

	return cleanedInput
}