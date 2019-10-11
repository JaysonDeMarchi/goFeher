package main

import (
    "github.com/JaysonDeMarchi/goFeher/battle"
    "github.com/JaysonDeMarchi/goFeher/unit"
    "github.com/JaysonDeMarchi/goFeher/special"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

func main() {
    anna := unit.Unit{
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

    axeFighter := unit.Unit{
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

    blueManakete := unit.Unit{
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

    battle.Battle(&anna, &axeFighter);
    battle.Battle(&anna, &blueManakete);

    db, err := sql.Open("mysql", "root:root_password@tcp(127.0.0.1:3306)/gofeher_database")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
}
