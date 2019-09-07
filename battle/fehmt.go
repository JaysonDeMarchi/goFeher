package battle

import (
    "fmt"
    "github.com/JaysonDeMarchi/goFeher/heroes"
)

func printIntro(){
    fmt.Println("---------------- Battle ----------------")
}

func printResults(battleSequence []Pair) {
    heroes.PrintStats(battleSequence[0].Attacker)
    heroes.PrintStats(battleSequence[0].Defender)
    fmt.Println("----------------------------------------\n")
}

func printSequence(battlePair Pair) {
    attacker := *battlePair.Attacker
    defender := *battlePair.Defender
    fmt.Printf(
        "%s (%d) -> %s (%d)\n",
        attacker.Name,
        attacker.CurrentHp,
        defender.Name,
        defender.CurrentHp)
}
