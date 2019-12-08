package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strconv"
)

// Game is the game struct which all the game loop is working for
type Game struct {
	WelcomeMessage     string
	Questions          []Question
	correctAnswerCount int
	numberOfQuestions  int
}

// Question is the question struct which is used by the game struct
type Question struct {
	QuestionText       string
	Answers            []string
	correctAnswerIndex int
}

func (g *Game) initiate() {
	fmt.Println(g.WelcomeMessage)
	g.generateQuestions()
}

func (g *Game) start() {
	g.correctAnswerCount = 0
	i := 0
	for i < len(g.Questions) {
		var givenAnswer int
		fmt.Println(g.Questions[i].QuestionText)
		a := 0
		for a < len(g.Questions[i].Answers) {
			fmt.Println(strconv.Itoa(a+1) + ") " + g.Questions[i].Answers[a])
			a++
		}
		fmt.Scan(&givenAnswer)
		if givenAnswer == g.Questions[i].correctAnswerIndex {
			fmt.Println("Answer is CORRECT!")
			g.correctAnswerCount++
		} else {
			fmt.Println("Answer is FALSE!")
		}
		fmt.Println(givenAnswer)
		i++
	}
	g.end()
}

func (g *Game) end() {
	fmt.Println("You have answered " + strconv.Itoa(g.correctAnswerCount) + " questions correctly!")
	fmt.Println()
}

func (g *Game) exit() {
	fmt.Println("Good exercise! See you soon!")
	fmt.Println()
	os.Exit(0)
}

func (g *Game) generateQuestions() {
	plan, _ := ioutil.ReadFile("dictionary.json")
	var data map[string]interface{}
	json.Unmarshal(plan, &data)

	keys := reflect.ValueOf(data).MapKeys()

	i := 0
	for i < g.numberOfQuestions {
		word := keys[rand.Intn(len(keys))].Interface().(string)
		var answersList []string

		correctAnswerIndex := rand.Intn(2)
		a := 0
		for a < 3 {
			if a == correctAnswerIndex {
				answersList = append(answersList, data[word].(string))
			} else {
				a := keys[rand.Intn(len(keys))].Interface().(string)
				answersList = append(answersList, data[a].(string))
			}
			a++
		}

		questionToAdd := Question{QuestionText: word, Answers: answersList, correctAnswerIndex: correctAnswerIndex + 1}
		g.Questions = append(g.Questions, questionToAdd)
		i++
	}
}

func main() {
	var choice int

	game := &Game{"Welcome to vocability. Do you want to start? \n1) Yes\n2) No", []Question{}, 0, 10}
	for {
		game.initiate()
		fmt.Scan(&choice)
		if choice == 1 {
			game.start()
		} else {
			game.exit()
		}
	}
}
