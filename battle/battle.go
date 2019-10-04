package battle

import (
    "github.com/JaysonDeMarchi/goFeher/unit"
)

const loggerIsActive bool = true

type Pair struct {
    Attacker, Defender *unit.Unit
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
    if attacker.Special.GetEffect().GetTrigger() == "attack" {
        if attacker.Special.GetCurrentCooldown() == 0 {
            attacker.Special.SetCurrentCooldown(attacker.Special.GetBaseCooldown())
        } else {
            attacker.Special.SetCurrentCooldown(attacker.Special.GetCurrentCooldown() - 1)
        }
    }
    if defender.Special.GetEffect().GetTrigger() == "attack" {
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
    if attacker.Special.GetEffect().GetTrigger() == "attack" && attacker.Special.GetCurrentCooldown() == 0 {
        trueDamage = attacker.Special.Trigger(attacker.Atk)
    }
    return trueDamage - (defender.Bulk(attacker))
}
