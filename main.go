package main

import (
	"fmt"
	"io"
	"os"
)

func mn() {
	game := NewQuizGame(os.Stdout, os.Stdin)
	err := game.Run()
	if err != nil {
		fmt.Println(err)
	}
}

type QuizGame struct {
	output io.Writer
	input  io.Reader
}

func NewQuizGame(output io.Writer, input io.Reader) *QuizGame {
	return &QuizGame{
		output: output,
		input:  input,
	}
}

func (q *QuizGame) Run() error {
	fmt.Fprintln(q.output, "Welcome to my riddle game!")

	player, err := q.askName()
	if err != nil {
		return err
	}

	if !q.verifyAge(player.Age) {
		fmt.Fprintln(q.output, "Hmm you don't seem old enough to play. I'll come back when you are...more ripe ")
		return nil // Exit after age check
	}

	fmt.Fprintln(q.output, fmt.Sprintf("Welcome to the game %v, I hope you are ready for what is to come\n", player.Name))

	if err := q.playRiddle("What is at the end of a rainbow?", "W"); err != nil {
		return err
	}

	if err := q.playRiddle("What kind of goose fights with snakes?", "A mongoose"); err != nil {
		return err
	}

	if err := q.playRiddle("What has an eye but can not see?", "A needle"); err != nil {
		return err
	}

	fmt.Fprintln(q.output, "Congratulations! You won the game.")
	return nil
}

type player struct {
	Name string
	Age  int
}

func (q *QuizGame) askName() (player, error) {
	fmt.Fprint(q.output, "What's your name? ")
	var name string
	_, err := fmt.Fscan(q.input, &name)
	return player{Name: name}, err
}

func (q *QuizGame) verifyAge(age int) bool {
	fmt.Fprint(q.output, "What's your age? ")
	_, err := fmt.Fscan(q.input, &age)
	if err != nil {
		return false
	}
	return age >= 10
}

func (q *QuizGame) playRiddle(riddle, answer string) error {
	fmt.Fprintln(q.output, fmt.Sprintf("\nHere is the riddle: %v", riddle))
	var userAnswer string
	_, err := fmt.Fscan(q.input, &userAnswer)
	if err != nil {
		return err
	}

	if userAnswer != answer {
		fmt.Fprintln(q.output, "Incorrect answer. :( Better luck next time.")
		return fmt.Errorf("incorrect answer")
	}

	fmt.Fprintln(q.output, "Correct! Keep going.")
	return nil
}

//string for words
//int for numbers
//uint for whole number (no: -6)
//float64 this is for double numbers in decimal places (2.5)
//bool can whether the value is true or false

// riddles and their answers
// que: What is at the end of a rainbow?
// ans: W
// que: What kind of goose fights with snakes?
// ans: A mongoose.
// que: When is music like vegetables?
// ans: When there are two beats (beets) to the measure.
