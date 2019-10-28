package main

import (
    "github.com/JaysonDeMarchi/goFeher/battle"
    "github.com/JaysonDeMarchi/goFeher/unit"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

func main() {
    anna := unit.New("anna");
    axeFighter := unit.New("axeFighter");
    blueManakete := unit.New("blueManakete");

    battle.Battle(anna, *axeFighter);
    battle.Battle(anna, *blueManakete);

    db, err := sql.Open("mysql", "root:root_password@tcp(127.0.0.1:3306)/gofeher_database")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
}
