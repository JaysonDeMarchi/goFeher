package unit

import (
    "github.com/JaysonDeMarchi/goFeher/special"
)

type Unit interface {
    Bulk(*Unit) int
    GetAssist() string
    GetAtk() int
    GetBaseHp() int
    GetCurrentHp() int
    SetCurrentHp(int)
    GetDef() int
    GetMovementType() string
    GetName() string
    GetRes() int
    GetSpd()int
    GetSpecial() special.Special
    GetWeapon() string
    GetWeaponColor() string
    GetWeaponType() string
}
