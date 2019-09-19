package battle

import (
    "github.com/JaysonDeMarchi/goFeher/heroes"
)

const loggerIsActive bool = true

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
    if attacker.Special.GetTriggerType() == "attack" {
        if attacker.Special.GetCurrentCooldown() == 0 {
            attacker.Special.SetCurrentCooldown(attacker.Special.GetBaseCooldown())
        } else {
            attacker.Special.SetCurrentCooldown(attacker.Special.GetCurrentCooldown() - 1)
        }
    }
    if defender.Special.GetTriggerType() == "attack" {
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

func Battle(attacker *heroes.Unit, defender *heroes.Unit) {
    battlePair := Pair{
        Attacker: attacker,
        Defender: defender,
    }
    battleSequence := buildBattleSequence(battlePair)
    if loggerIsActive {
        printIntro()
    }
    for _, battlePair := range battleSequence {
        attack(battlePair)
        if loggerIsActive {
            printSequence(battlePair)
        }
        if battlePair.Defender.CurrentHp == 0 {
            break
        }
    }

    if loggerIsActive {
        printResults(battleSequence)
    }
}

func calculateDamage(battlePair Pair) int {
    attacker, defender := battlePair.Attacker, battlePair.Defender
    trueDamage := attacker.Atk
    if attacker.Special.GetTriggerType() == "attack" && attacker.Special.GetCurrentCooldown() == 0 {
        trueDamage = attacker.Special.Trigger(attacker.Atk)
    }
    return trueDamage - (defender.Bulk(attacker))
}
