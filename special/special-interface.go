package special

type Special interface {
    Trigger(int) int
    GetBaseCooldown() int
    GetCurrentCooldown() int
    SetCurrentCooldown(int)
    GetEffectValue() float64
    GetName() string
    GetTriggerType() string
    GetType() string
}
