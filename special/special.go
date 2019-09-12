package special

type special struct {
    BaseCooldown, CurrentCooldown int
    EffectValue float64
    Name, Trigger, Type string
}

func New(name string) Special {
    return &special{
        Name: "night sky",
        BaseCooldown: 3,
        CurrentCooldown: 3,
        EffectValue: 0.5,
        Trigger: "attack",
        Type: "damage boost mult",
    }
}

func (s *special) GetBaseCooldown() int {
    return s.BaseCooldown
}

func (s *special) GetCurrentCooldown() int {
    return s.CurrentCooldown
}

func (s *special) SetCurrentCooldown(currentCooldown int) {
    s.CurrentCooldown = currentCooldown
}

func (s *special) GetEffectValue() float64 {
    return s.EffectValue
}

func (s *special) GetName() string {
    return s.Name
}

func (s *special) GetTrigger() string {
    return s.Trigger
}

func (s *special) GetType() string {
    return s.Type
}
