package battle

import (
    "github.com/JaysonDeMarchi/goFeher/heroes"
)

type Pair struct {
    Attacker, Defender *heroes.Unit
}

func calculateRemainingHp(battlePair Pair) int {
    defender := battlePair.Defender
    damage := calculateDamage(battlePair)
    if defender.CurrentHp < damage {
        return 0
    }
    return defender.CurrentHp - damage
}


func attack(battlePair Pair) {
    battlePair.Defender.SetCurrentHp(calculateRemainingHp(battlePair))
}

func buildBattleSequence(battlePair Pair) []Pair {
    attacker := battlePair.Attacker
    defender := battlePair.Defender
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

func Battle(attacker *heroes.Unit, defender *heroes.Unit, showResults bool) {
    battlePair := Pair{
        Attacker: attacker,
        Defender: defender,
    }
    battleSequence := buildBattleSequence(battlePair)

    if showResults {
        printIntro()
    }

    for _, battlePair := range battleSequence {
        attack(battlePair)
        if showResults {
            printSequence(battlePair)
        }
    }

    if showResults {
        printResults(battleSequence)
    }
}

func calculateDamage(battlePair Pair) int {
    return battlePair.Attacker.Atk - (battlePair.Defender.Bulk(battlePair.Attacker))
}
