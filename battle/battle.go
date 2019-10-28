package battle

import (
    "github.com/JaysonDeMarchi/goFeher/unit"
)

const loggerIsActive bool = true

type Pair struct {
    Attacker, Defender *unit.Unit
}

func calculateRemainingHp(battlePair Pair) int {
    defender := *battlePair.Defender
    damage := calculateDamage(battlePair)
    if defender.GetCurrentHp() < damage {
        return 0
    }
    return defender.GetCurrentHp() - damage
}

func updateSpecialCooldowns(battlePair Pair) {
    attacker, defender := *battlePair.Attacker, *battlePair.Defender
    if attacker.GetSpecial().GetEffect().GetTrigger() == "attack" {
        if attacker.GetSpecial().GetCurrentCooldown() == 0 {
            attacker.GetSpecial().SetCurrentCooldown(attacker.GetSpecial().GetBaseCooldown())
        } else {
            attacker.GetSpecial().SetCurrentCooldown(attacker.GetSpecial().GetCurrentCooldown() - 1)
        }
    }
    if defender.GetSpecial().GetEffect().GetTrigger() == "attack" {
        defender.GetSpecial().SetCurrentCooldown(defender.GetSpecial().GetCurrentCooldown() - 1)
    }
}

func attack(battlePair Pair) {
    defender := *battlePair.Defender
    defender.SetCurrentHp(calculateRemainingHp(battlePair))
    updateSpecialCooldowns(battlePair)
}

func buildBattleSequence(battlePair Pair) []Pair {
    attacker, defender := *battlePair.Attacker, *battlePair.Defender
    battleSequence := []Pair {
        Pair{
            battlePair.Attacker,
            battlePair.Defender,
        },
        Pair{
            battlePair.Defender,
            battlePair.Attacker,
        },
    }
    if attacker.GetSpd() - defender.GetSpd() >= 5 {
        battleSequence = append(battleSequence, Pair{
            battlePair.Attacker,
            battlePair.Defender,
        })
    } else if defender.GetSpd() -  attacker.GetSpd() >= 5 {
        battleSequence = append(battleSequence, Pair{
            battlePair.Defender,
            battlePair.Attacker,
        })
    }
    return battleSequence
}

func Battle(attacker *unit.Unit, defender *unit.Unit) {
    battlePair := Pair{
        Attacker: attacker,
        Defender: defender,
    }
    battleSequence := buildBattleSequence(battlePair)
    if loggerIsActive {
        printIntro()
    }
    for _, battlePair := range battleSequence {
        defender := *battlePair.Defender
        attack(battlePair)
        if loggerIsActive {
            printSequence(battlePair)
        }
        if defender.GetCurrentHp() == 0 {
            break
        }
    }

    if loggerIsActive {
        printResults(battleSequence)
    }
}

func calculateDamage(battlePair Pair) int {
    attacker := *battlePair.Attacker
    defender := *battlePair.Defender
    trueDamage := attacker.GetAtk()
    if attacker.GetSpecial().GetEffect().GetTrigger() == "attack" && attacker.GetSpecial().GetCurrentCooldown() == 0 {
        trueDamage = attacker.GetSpecial().Trigger(attacker.GetAtk())
    }
    return trueDamage - (defender.Bulk(battlePair.Attacker))
}
