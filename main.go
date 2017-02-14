package main

import (
	"bufio"
	"fmt"
	"os"	
)

type Transaction struct {
	property int
}

type Reward struct {
	property int
}

type Challenge struct {
	property int
}

type User struct {
	username string
	ledger []Transaction
	balance int
	rewards map[string]Reward
	challenge map[string]Challenge
}

type Game struct {
	user User
}

func (g Game) Init(u string) {
	fmt.Println("initializing game...")
	g.user.username = u
}

func (g Game) Command(c string) {
	fmt.Println("you typed: ", c)
}

func main() {
	var g Game
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("q to quit. enter name:")
	s.Scan()
	g.Init(s.Text())
	fmt.Print("command: ")
	for s.Scan() {
		input := s.Text()
		if input == "q" {
			break
		}
		g.Command(input)
		fmt.Print("command: ")
	}
}