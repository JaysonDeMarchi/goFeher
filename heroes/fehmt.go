package heroes

import (
    "fmt"
)

func PrintStats(unit *Unit) {
    fmt.Printf("NAME: %s\n", unit.Name)
    fmt.Printf("\tWEAPON TYPE: %s %s\n", unit.WeaponColor, unit.WeaponType)
    fmt.Printf("\tHP: %d / %d\t\tWEAPON: %s\n", unit.CurrentHp, unit.BaseHp, unit.Weapon)
    fmt.Printf("\tATK: %d\tSPD: %d\tASSIST: %s\n", unit.Atk, unit.Spd, unit.Assist)
    fmt.Printf("\tDEF: %d\tRES: %d\tSPECIAL: %s\n", unit.Def, unit.Res, unit.Special)
}
