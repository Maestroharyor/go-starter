package main

import (
	"bufio"
	"fmt"
	"os"
	"quizza/data"
	"strings"
)

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	// Variables
	var name string
	var score int

	// Welcome the user
	fmt.Println("Welcome to Quizza")

	// Get the user's name
	name = getUserInput("Please enter your name to continue: ")
	fmt.Printf("Hello %s! Let's play Quizza. \n\n", name)

	fmt.Println("Note you can type stop to stop the game at any time.")

	for i, quiz := range data.Questions {
		// fmt.Printf("Question #%d: %v \n\n", i+1, quiz.Question)
		formattedQuestion := fmt.Sprintf("Question #%d: %v \n", i+1, quiz.Question)
		answer := getUserInput(formattedQuestion)

		if answer == "stop" {
			break
		}
		if strings.EqualFold(answer, quiz.Answer) {
			fmt.Println("\nYou got it right!")
			score += int(quiz.Points)
			fmt.Printf("Your current score is: %v \n\n\n", score)
		} else {
			fmt.Println("\nYou got it wrong!")
			fmt.Println("The correct answer is: ", quiz.Answer)
			fmt.Printf("Your current score is: %v \n\n\n", score)
		}
	}

	finalPercent := (float64(score) / float64(len(data.Questions)*10)) * 100
	fmt.Printf("Thanks for playing Quizza %v! Your final score is: %v \nYour final percentage is: %.2f%% \n\n", name, score, finalPercent)

}
