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

func updateSpecialCooldowns(battlePair Pair) {
    attacker, defender := battlePair.Attacker, battlePair.Defender
    if attacker.Special.GetBaseCooldown() > 0 {
        attacker.Special.SetCurrentCooldown(attacker.Special.GetCurrentCooldown() - 1)
    }
    if defender.Special.GetBaseCooldown() > 0 {
        defender.Special.SetCurrentCooldown(defender.Special.GetCurrentCooldown() - 1)
    }
}

func attack(battlePair Pair) {
    battlePair.Defender.SetCurrentHp(calculateRemainingHp(battlePair))
    updateSpecialCooldowns(battlePair)
}

func buildBattleSequence(battlePair Pair) []Pair {
    attacker, defender := battlePair.Attacker, battlePair.Defender
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
        if battlePair.Defender.CurrentHp == 0 {
            break
        }
    }

    if showResults {
        printResults(battleSequence)
    }
}

func calculateDamage(battlePair Pair) int {
    attacker, defender := battlePair.Attacker, battlePair.Defender
    return attacker.Atk - (defender.Bulk(attacker))
}
