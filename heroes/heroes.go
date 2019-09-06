package heroes

type Unit struct {
    BaseHp, CurrentHp, Atk, Spd, Def, Res int
    Name, WeaponColor, WeaponType, MovementType, Weapon, Assist, Special string
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
