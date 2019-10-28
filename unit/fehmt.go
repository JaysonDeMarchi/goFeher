package unit

import (
    "fmt"
)

func PrintStats(unitPointer *Unit) {
    unit := *unitPointer
    fmt.Printf("NAME: %s\n",
        unit.GetName())
    fmt.Printf("\tWEAPON TYPE: %s %s\n",
        unit.GetWeaponColor(),
        unit.GetWeaponType())
    fmt.Printf("\tHP: %d / %d\t\tWEAPON: %s\n",
        unit.GetCurrentHp(),
        unit.GetBaseHp(),
        unit.GetWeapon())
    fmt.Printf("\tATK: %d\tSPD: %d\t\tASSIST: %s\n",
        unit.GetAtk(),
        unit.GetSpd(),
        unit.GetAssist())
    fmt.Printf("\tDEF: %d\tRES: %d\t\tSPECIAL: %s (%d / %d)\n",
        unit.GetDef(),
        unit.GetRes(),
        unit.GetSpecial().GetName(),
        unit.GetSpecial().GetCurrentCooldown(),
        unit.GetSpecial().GetBaseCooldown())
}
