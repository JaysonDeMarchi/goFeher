package main

import (
    "fmt"
)

type Unit struct {
    BaseHp, CurrentHp, Atk, Spd, Def, Res int
    Name, WeaponColor, WeaponType, MovementType, Weapon, Assist, Special string
}

type BattlePair struct {
    Attacker, Defender Unit
}

func attack(attacker Unit, defender Unit) int {
    damage := calculateDamage(attacker, defender)
    if defender.CurrentHp < damage {
        return 0
    }
    return defender.CurrentHp - damage
}

func buildBattleSequence(attacker Unit, defender Unit) []BattlePair {

    defender.CurrentHp = attack(attacker, defender)
    attacker.CurrentHp = attack(defender, attacker)
    if attacker.Spd - defender.Spd >= 5 {
        defender.CurrentHp = attack(attacker, defender)
    }
    if defender.Spd -  attacker.Spd >= 5 {
        attacker.CurrentHp = attack(defender, attacker)
    }
}

func battle(attacker Unit, defender Unit, showResults bool) {
    battleSequence := buildBattleSequence(attacker, defender)

    for battlePair := range battleSequence {
        battlePair.Defender.currentHp = attack(battlePair.Attacker, battlePair.Defender)
    }

    if showResults {
        results(battleSequence)
    }
}

func printSequence(battleSequence []BattlePair) {
    for attacker, defender := range battleSequence {
        printf("%s -> %s\n", attacker, defender)
    }
}

func printStats(unit Unit) {
    fmt.Printf("NAME: %s\n", unit.Name)
    fmt.Printf("\tWEAPON TYPE: %s %s\n", unit.WeaponColor, unit.WeaponType)
    fmt.Printf("\tHP: %d / %d\t\tWEAPON: %s\n", unit.CurrentHp, unit.BaseHp, unit.Weapon)
    fmt.Printf("\tATK: %d\tSPD: %d\tASSIST: %s\n", unit.Atk, unit.Spd, unit.Assist)
    fmt.Printf("\tDEF: %d\tRES: %d\tSPECIAL: %s\n", unit.Def, unit.Res, unit.Special)
}

func results(battleSequence) {
    fmt.Println("--- Battle ---")
    printSequence(battleSequence)
    printStats(battleSequence.Attacker)
    printStats(battleSequence.Defender)
    fmt.Println("--------------\n")
}

func (unit *Unit) Bulk(attacker Unit) int {
        magicalWeapons := map[string]bool {
            "tome": true,
            "dragonstone": true,
            "staff": true,
        }
        if magicalWeapons[attacker.WeaponType] {
            return unit.Res
        }
        return unit.Def
}

func calculateDamage(attacker Unit, defender Unit) int {
    return attacker.Atk - (defender.Bulk(attacker))
}

func main() {
    anna := Unit{
        22,
        22,
        22,
        16,
        7,
        9,
        "Anna",
        "green",
        "axe",
        "infantry",
        "silver axe",
        "",
        "night sky",
    }

    axeFighter := Unit{
        22,
        22,
        14,
        4,
        10,
        6,
        "Axe Fighter",
        "green",
        "axe",
        "infantry",
        "brave axe",
        "",
        "",
    }

    blueManakete := Unit{
        22,
        22,
        22,
        11,
        10,
        9,
        "Blue Manakete",
        "blue",
        "dragonstone",
        "infantry",
        "flametongue",
        "",
        "",
    }

    battle(anna, axeFighter, true);
    battle(anna, blueManakete, true);
}
