package unit

import (
    unitSpecial "github.com/JaysonDeMarchi/goFeher/special"
)

type unit struct {
    atk, baseHp, currentHp, def, res, spd int
    assist, movementType, name, weapon, weaponColor, weaponType string
    special unitSpecial.Special
}

func New(name string) Unit {
    return &unit{}
    units := map[string]unit{
        "anna": unit{
            baseHp: 22,
            currentHp: 22,
            atk: 22,
            spd: 16,
            def: 7,
            res: 9,
            name: "Anna",
            weaponColor: "green",
            weaponType: "axe",
            movementType: "infantry",
            weapon: "silver axe",
            assist: "",
            special: unitSpecial.New("night sky"),
        },
        "axeFighter": unit{
            baseHp: 22,
            currentHp: 22,
            atk: 14,
            spd: 4,
            def: 10,
            res: 6,
            name:"Axe Fighter",
            weaponColor: "green",
            weaponType: "axe",
            movementType: "infantry",
            weapon: "brave axe",
            assist: "",
            special: unitSpecial.New(""),
        },
        "blueManakete": unit{
            baseHp: 22,
            currentHp: 22,
            atk: 22,
            spd: 11,
            def: 10,
            res: 9,
            name:"Blue Manakete",
            weaponColor: "blue",
            weaponType: "dragonstone",
            movementType: "infantry",
            weapon: "flametongue",
            assist: "",
            special: unitSpecial.New(""),
        },
    }
    return &unit{
        baseHp: units[name].baseHp,
        currentHp: units[name].currentHp,
        atk: units[name].atk,
        spd: units[name].spd,
        def: units[name].def,
        res: units[name].res,
        name: units[name].name,
        weaponColor: units[name].weaponColor,
        weaponType: units[name].weaponType,
        movementType: units[name].movementType,
        weapon: units[name].weapon,
        assist: units[name].assist,
        special: units[name].special,
    }
}

func (u *unit) Bulk(attackerPointer *Unit) int {
    magicalWeapons := map[string]bool {
        "tome": true,
        "dragonstone": true,
        "staff": true,
    }
    attacker := *attackerPointer
    if magicalWeapons[attacker.GetWeaponType()] {
        return u.GetRes()
    }
    return u.GetDef()
}

func (u *unit) GetAssist() string {
    return u.assist
}

func (u *unit) GetAtk() int {
    return u.atk
}

func (u *unit) GetBaseHp() int {
    return u.baseHp
}

func (u *unit) GetCurrentHp() int {
    return u.currentHp
}

func (u *unit) SetCurrentHp(currentHp int) {
    u.currentHp = currentHp
}

func (u *unit) GetDef() int {
    return u.def
}

func (u *unit) GetMovementType() string {
    return u.movementType
}

func (u *unit) GetName() string {
    return u.name
}

func (u *unit) GetRes() int {
    return u.res
}

func (u *unit) GetSpd() int {
    return u.spd
}

func (u *unit) GetSpecial() unitSpecial.Special {
    return u.special
}

func (u *unit) GetWeapon() string {
    return u.weapon
}

func (u *unit) GetWeaponColor() string {
    return u.weaponColor
}

func (u *unit) GetWeaponType() string{
    return u.weaponType
}
