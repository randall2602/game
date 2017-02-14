package main

import (
	"bufio"
	"fmt"
	"os"
)

const HELP_TEXT = "\n" +
"game is a tool for managing goals.\n" +
"\n" +
"commands:\n" +
"\tadd\tadd a reward or challenge\n" +
"\tremove\tremove a reward or challenge\n" +
"\tlist\tlist rewards, challenges, or ledger entries\n" +
"\tedit\tedit a reward, challenge, or ledger entry\n" +
"\tq\tquit"

type transactionType int

const (
	Credit transactionType = iota
	Debit
)

type timescale int

const (
	Day timescale = iota
	Week
	Month
	Year
)

type Transaction struct {
	entryType transactionType
	itemType string
	amount int
	points int
	description string
}

type Reward struct {
	weight float64
	unit string
	amount int
	timescale timescale
}

type Challenge struct {
	weight float64
	unit string
	amount int
	timescale timescale
}

type User struct {
	username string
	ledger []Transaction
	balance int
}

type Game struct {
	user User
	rewards map[string]Reward
	challenge map[string]Challenge
}

func (g Game) Init(u string) {
	fmt.Println("initializing game...")
	g.user.username = u
	g.user.balance = 1000000
}

func (g Game) Command(c string, s *bufio.Scanner) {
	switch c {
	case "add":
		g.Add(s)
	//case "remove":
	//	g.Remove(s)
	//case "list":
	//	g.List(s)
	//case "edit":
	//	g.Edit(s)
	case "help":
		g.Help()
	default:
		fmt.Println("command not recognized...")
		g.Help()
	}
}

func (g Game) Add(s *bufio.Scanner) {
	fmt.Println("add a reward or challenge?")
	s.Scan()
	input := s.Text()
	switch input {
	case "reward":
		fmt.Println("you chose to add a reward")
	case "challenge":
		fmt.Println("you chose to add a challenge")
	default:
		fmt.Println("type \"reward\" or \"challenge\".")
		g.Add(s)
	}
}

func (g Game) Help() {
	fmt.Println(HELP_TEXT)
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
		g.Command(input, s)
		fmt.Print("command: ")
	}
}
