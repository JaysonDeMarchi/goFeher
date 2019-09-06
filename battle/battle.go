package battle

import (
    "github.com/JaysonDeMarchi/goFeher/heroes"
)

type Pair struct {
    Attacker, Defender heroes.Unit
}

func attack(attacker heroes.Unit, defender heroes.Unit) int {
    damage := calculateDamage(attacker, defender)
    if defender.CurrentHp < damage {
        return 0
    }
    return defender.CurrentHp - damage
}

func buildBattleSequence(attacker heroes.Unit, defender heroes.Unit) []Pair {
    battleSequence := []Pair {
        Pair{
            attacker,
            defender,
        },
        Pair{
            defender,
            attacker,
        },
    }
    if attacker.Spd - defender.Spd >= 5 {
        battleSequence = append(battleSequence, Pair{
            attacker,
            defender,
        })
    } else if defender.Spd -  attacker.Spd >= 5 {
        battleSequence = append(battleSequence, Pair{
            defender,
            attacker,
        })
    }
    return battleSequence
}

func Battle(attacker heroes.Unit, defender heroes.Unit, showResults bool) {
    battleSequence := buildBattleSequence(attacker, defender)

    if showResults {
        printIntro()
    }

    for _, battlePair := range battleSequence {
        battlePair.Defender.CurrentHp = attack(battlePair.Attacker, battlePair.Defender)
        if showResults {
            printSequence(battlePair)
        }
    }

    if showResults {
        printResults(battleSequence)
    }
}

func calculateDamage(attacker heroes.Unit, defender heroes.Unit) int {
    return attacker.Atk - (defender.Bulk(attacker))
}
