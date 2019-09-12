package main

import (
    "github.com/JaysonDeMarchi/goFeher/battle"
    "github.com/JaysonDeMarchi/goFeher/heroes"
    "github.com/JaysonDeMarchi/goFeher/special"
)

func main() {
    anna := heroes.Unit{
        BaseHp: 22,
        CurrentHp: 22,
        Atk: 22,
        Spd: 16,
        Def: 7,
        Res: 9,
        Name: "Anna",
        WeaponColor: "green",
        WeaponType: "axe",
        MovementType: "infantry",
        Weapon: "silver axe",
        Assist: "",
        Special: special.New("night sky"),
    }

    axeFighter := heroes.Unit{
        BaseHp: 22,
        CurrentHp: 22,
        Atk: 14,
        Spd: 4,
        Def: 10,
        Res: 6,
        Name:"Axe Fighter",
        WeaponColor: "green",
        WeaponType: "axe",
        MovementType: "infantry",
        Weapon: "brave axe",
        Assist: "",
        Special: special.New(""),
    }

    blueManakete := heroes.Unit{
        BaseHp: 22,
        CurrentHp: 22,
        Atk: 22,
        Spd: 11,
        Def: 10,
        Res: 9,
        Name:"Blue Manakete",
        WeaponColor: "blue",
        WeaponType: "dragonstone",
        MovementType: "infantry",
        Weapon: "flametongue",
        Assist: "",
        Special: special.New(""),
    }

    battle.Battle(&anna, &axeFighter, true);
    battle.Battle(&anna, &blueManakete, true);
}
