package special

import (
    "github.com/JaysonDeMarchi/goFeher/special/effect"
)

type Special interface {
    Trigger(int) int
    GetBaseCooldown() int
    GetCurrentCooldown() int
    SetCurrentCooldown(int)
    GetEffect() effect.Effect
    GetName() string
}
