package battle

import (
    "fmt"
    "github.com/JaysonDeMarchi/goFeher/unit"
)

func printIntro(){
    fmt.Println("---------------- Battle ----------------")
}

func printResults(battleSequence []Pair) {
    unit.PrintStats(battleSequence[0].Attacker)
    unit.PrintStats(battleSequence[0].Defender)
    fmt.Println("----------------------------------------\n")
}

func printSequence(battlePair Pair) {
    attacker := *battlePair.Attacker
    defender := *battlePair.Defender
    fmt.Printf(
        "%s (%d) -> %s (%d)\n",
        attacker.GetName(),
        attacker.GetCurrentHp(),
        defender.GetName(),
        defender.GetCurrentHp())
}
