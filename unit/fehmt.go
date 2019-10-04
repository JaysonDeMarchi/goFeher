package unit

import (
    "fmt"
)

func PrintStats(unit *Unit) {
    fmt.Printf("NAME: %s\n",
        unit.Name)
    fmt.Printf("\tWEAPON TYPE: %s %s\n",
        unit.WeaponColor,
        unit.WeaponType)
    fmt.Printf("\tHP: %d / %d\t\tWEAPON: %s\n",
        unit.CurrentHp,
        unit.BaseHp,
        unit.Weapon)
    fmt.Printf("\tATK: %d\tSPD: %d\t\tASSIST: %s\n",
        unit.Atk,
        unit.Spd,
        unit.Assist)
    fmt.Printf("\tDEF: %d\tRES: %d\t\tSPECIAL: %s (%d / %d)\n",
        unit.Def,
        unit.Res,
        unit.Special.GetName(),
        unit.Special.GetCurrentCooldown(),
        unit.Special.GetBaseCooldown())
}
