package actions

import (
	"math/rand"
	"time"
)

var currentPlayerHealth = PLAYER_HEALTH
var currentMonsterHealth = MONSTER_HEALTH
var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG
	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dmgValue
	return dmgValue
}
func HealPlayer() int {
	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)
	healthDiff := PLAYER_HEALTH - currentPlayerHealth
	if healthDiff >= healValue {
		currentPlayerHealth += healValue
		return healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
		return healthDiff
	}
}
func AttackPlayer() int {
	dmgValue := generateRandBetween(MONSTER_ATTACK_MIN_DMG, MONSTER_ATTACK_MAX_DMG)
	currentPlayerHealth -= dmgValue
	return dmgValue

}
func GetHealthAmount() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}
func generateRandBetween(min, max int) int {
	return randGenerator.Intn(max-min) + min
}
