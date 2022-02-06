package main

import (
	"fmt"
)

const greeting = "Hallo en welkom bij de wb Textadventure!"

var playername string
var playergender string
var playerApp string
var playerAge uint
var playerChoice string

func main() {

	fmt.Println(greeting)
	playerInfoFunc(playername, playergender, playerApp, playerAge)

}

func playerInfoFunc(playername string, playergender string, playerApp string, playerAge uint) {
	//Name
	fmt.Println("Wat is je naam? \n")
	fmt.Scan(&playername)
	fmt.Printf("Dus ja naam is %v \n", playername)
	//gender
	fmt.Printf("%v wat is je gender? \n", playername)
	fmt.Scan(&playergender)
	fmt.Printf("dus je bent een %v van gender? \n", playergender)
	//appearance
	fmt.Printf("%v Hoe zie er uit? \n", playername)
	fmt.Scan(&playerApp)
	fmt.Printf("dus je ziet er zo uit %v %v? \n", playerApp, playername)
	//age
	fmt.Printf("Wat is je leeftijd %v? \n", playername)
	fmt.Scan(&playerAge)
	fmt.Printf("dus %v je bent %v jaar oud? \n", playername, playerAge)

	//recap
	fmt.Printf("[Jari & Nick]Hallo daar %v we willen je graag welkom heten in de wereld van PenK \n", playername)
	fmt.Printf("Om het even samen te vatten: \n -------------------------------\n je naam is: %v \n je gender is: %v  \n je ziet er zo uit: %v \n en je bent: %v jaar oud? \n -------------------------------", playername, playergender, playerApp, playerAge)

}
