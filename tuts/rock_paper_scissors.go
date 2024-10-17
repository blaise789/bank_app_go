package main

import (
	"fmt"
	"math/rand"
	"strings"
)
func playRockPaperScissors(){
	rounds:=3
	for i:=0; i<=rounds;i++{
	compRandomNum:=rand.Intn(3)
	var computerChoice string
	switch compRandomNum{
		case 0:
            computerChoice="rock"
        case 1:
            computerChoice="paper"
        case 2:
            computerChoice="scissors"
        }
		fmt.Print("enter your choice(Rock,Paper,Scissors)")
		var usersChoice string
		fmt.Scanln(&usersChoice)
		usersChoice=strings.ToLower(usersChoice)     
		switch {
		case usersChoice==computerChoice:
			fmt.Println("game tie")
		case usersChoice=="paper" && computerChoice=="rock" ,usersChoice=="scissors" && computerChoice=="paper",usersChoice=="rock" && computerChoice=="scissors":
			fmt.Println(computerChoice)
			fmt.Println("you win")
	    default:
			fmt.Println("you lose")		
		}       
			//play game here
			// check for winner
			// display result
			// ask for another round
			// repeat until game over
			
      
	}
	
	}


