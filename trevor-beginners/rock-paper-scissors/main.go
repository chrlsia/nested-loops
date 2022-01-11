package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	clearScreen()

	for counter:=1;counter<=3;counter++{
		playerChoice := ""
		playerValue := -1

		computerValue := rand.Intn(2)

		reader := bufio.NewReader(os.Stdin)


		fmt.Print("Please enter rock, paper, or scissors -> ")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		}

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
			// break
		case PAPER:
			fmt.Println("computer chose PAPER")
			// break
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
		default:
		}

		// fmt.Println()

		if computerValue == playerValue {
			fmt.Println("It is a draw")
		} else {
			switch playerValue{
			case ROCK:
				if playerValue==PAPER{
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Play wins")
				}
				break
			case PAPER:
				if playerValue==SCISSORS{
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Play wins")
				}
				break
			case SCISSORS:
				if playerValue==ROCK{
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Play wins")
				}
				break
			default:
				fmt.Println("Invalid choice")
			}
		}
	}
}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
