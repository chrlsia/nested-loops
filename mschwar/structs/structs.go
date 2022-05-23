package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName   string
	lastName    string
	birthDate   string
	createdDate time.Time
}

func main() {
	var newUser User

	firstName := getUserData("Please enter your first name:")
	lastName := getUserData("Please enter your last name:")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY):")
	created := time.Now()

	newUser = User{
		firstName:   firstName,
		lastName:    lastName,
		birthDate:   birthdate,
		createdDate: created,
	}

	// ... do something awesome with these data
	fmt.Println(newUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')
	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}
