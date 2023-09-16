package main

import (
	"github.com/Alifarid0011/monster_game/actions"
	"github.com/Alifarid0011/monster_game/interaction"
)

var currentRound = 0

func main() {
	startGame()
	winner := "" //"Player || "Monster
	for winner == "" {
		winner = executeRound()
	}
	endGame(winner)
}
func startGame() {
	interaction.PrintGreeting()
}
func executeRound() string {
	currentRound++
	isisSpecialRound := currentRound%3 == 0
	interaction.ShowAvailableActions(isisSpecialRound)
	userChoice := interaction.GetPlayerChoice(isisSpecialRound)
	var playerHealth int
	var monsterHealth int
	var playerDMG int
	var playerHealVAL int
	var monsterDMG int
	if userChoice == "ATTACK" {
		playerDMG = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealVAL = actions.HealPlayer()
	} else {
		playerDMG = actions.AttackMonster(true)
	}
	monsterDMG = actions.AttackPlayer()
	playerHealth, monsterHealth = actions.GetHealthAmount()
	NewRoundData := interaction.NewRound(userChoice, playerHealth, playerHealVAL, playerDMG, monsterHealth, monsterDMG)
	interaction.PrintRoundStatistics(NewRoundData)
	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}
	return ""
}
func endGame(winner string) {
	interaction.DeclareWinner(winner)
}
