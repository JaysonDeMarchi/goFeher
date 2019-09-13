package special

type special struct {
    BaseCooldown, CurrentCooldown int
    EffectValue float64
    Name, TriggerType, Type string
}

func New(name string) Special {
    if name == "" {
        return &special{}
    }
    specials := map[string]special {
        "night sky": special{
            Name: "night sky",
            BaseCooldown: 3,
            CurrentCooldown: 3,
            EffectValue: 0.5,
            TriggerType: "attack",
            Type: "damage boost mult",
        },
    }
    return &special{
        Name: specials[name].Name,
        BaseCooldown: specials[name].BaseCooldown,
        CurrentCooldown: specials[name].CurrentCooldown,
        EffectValue: specials[name].EffectValue,
        TriggerType: specials[name].TriggerType,
        Type: specials[name].Type,
    }
}
func (s *special) Trigger(base int) int {
    if s.Type == "damage boost mult" {
        return int(float64(base) * s.EffectValue)
    }
    return base
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

func (s *special) GetTriggerType() string {
    return s.TriggerType
}

func (s *special) GetType() string {
    return s.Type
}
