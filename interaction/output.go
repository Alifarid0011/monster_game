package interaction

import "fmt"

type Round struct {
	Action         string
	PlayerHeath    int
	PlayerHeathVal int
	PlayerDMG      int
	MonsterHeath   int
	MonsterDMG     int
}

func NewRound(
	Action string,
	PlayerHeath,
	PlayerHeathVal,
	PlayerDMG,
	MonsterHeath,
	MonsterDMG int,
) (NewRd *Round) {
	NewRd = &Round{
		Action,
		PlayerHeath,
		PlayerHeathVal,
		PlayerDMG,
		MonsterHeath,
		MonsterDMG,
	}
	return
}
func PrintGreeting() {
	fmt.Println("Monster slayer")
	fmt.Println("starting a new game ...")
	fmt.Println("Good luck !!")
}
func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("__________________________")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal ")
	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}
func DeclareWinner(winner string) {
	fmt.Println("------------------------")
	fmt.Println("Game Over !")
	fmt.Println("------------------------")
	fmt.Printf("%v won !", winner)

}
func PrintRoundStatistics(round *Round) {
	if round.Action == "ATTACK" {
		fmt.Printf("Player attacker monster for %v damage. \n", round.PlayerDMG)
	} else if round.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster  for %v damage. \n", round.PlayerDMG)
	} else {
		fmt.Printf("Player healed for %v. \n", round.PlayerHeathVal)
	}
	fmt.Printf("Monster attacked player for %v damage. \n", round.MonsterDMG)
	fmt.Printf("Player heath :%v \n", round.PlayerHeath)
	fmt.Printf("Monster heath :%v \n", round.MonsterHeath)
}
