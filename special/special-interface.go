package special

type Special interface {
    GetBaseCooldown() int
    GetCurrentCooldown() int
    SetCurrentCooldown(int)
    GetEffectValue() float64
    GetName() string
    GetTrigger() string
    GetType() string
}
