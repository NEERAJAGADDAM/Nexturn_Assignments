package main

import (
	"fmt"
	"strconv"
	"time"
)

// Struct to represent a question
type Question struct {
	QuestionText  string
	Options       [4]string
	CorrectAnswer int
}

func main() {
	// Question Bank
	questions := []Question{
		{
			QuestionText:  "What is the capital of France?",
			Options:       [4]string{"1. Paris", "2. Berlin", "3. Madrid", "4. Rome"},
			CorrectAnswer: 1,
		},
		{
			QuestionText:  "What is the largest planet in our solar system?",
			Options:       [4]string{"1. Earth", "2. Jupiter", "3. Saturn", "4. Mars"},
			CorrectAnswer: 2,
		},
		{
			QuestionText:  "Who developed the theory of relativity?",
			Options:       [4]string{"1. Isaac Newton", "2. Albert Einstein", "3. Galileo Galilei", "4. Nikola Tesla"},
			CorrectAnswer: 2,
		},
	}

	var score int
	//var userAnswer int
	var userInput string
	timer := 10 * time.Second // Time limit for each question

	// Take the quiz
	for i, q := range questions {
		// Display question and options
		fmt.Println("Question", i+1, ":", q.QuestionText)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		// Timer starts
		timerExpired := make(chan bool)
		go func() {
			time.Sleep(timer)
			timerExpired <- true
		}()

		// User input with error handling
		inputDone := make(chan bool)
		go func() {
			fmt.Print("Enter your answer (1-4) or 'exit' to quit: ")
			fmt.Scanln(&userInput)

			// Handle exit command
			if userInput == "exit" {
				fmt.Println("You exited the quiz.")
				close(inputDone)
				return
			}

			userAnswer, err := strconv.Atoi(userInput)
			if err != nil || userAnswer < 1 || userAnswer > 4 {
				fmt.Println("Invalid input, please enter a number between 1 and 4.")
			} else {
				// Valid input, process the answer
				if userAnswer == q.CorrectAnswer {
					score++
				}
				close(inputDone)
			}
		}()

		// Check if timer expired or user input is done
		select {
		case <-timerExpired:
			fmt.Println("Time's up for this question!")
			close(inputDone)
		case <-inputDone:
		}
	}

	// Score Calculation
	fmt.Println("\nYour final score is:", score, "/", len(questions))

	// Performance classification
	switch {
	case score == len(questions):
		fmt.Println("Excellent!")
	case score >= len(questions)/2:
		fmt.Println("Good!")
	default:
		fmt.Println("Needs Improvement!")
	}
}
