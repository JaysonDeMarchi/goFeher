package special

import (
    "github.com/JaysonDeMarchi/goFeher/special/effect"
)

type special struct {
    BaseCooldown, CurrentCooldown int
    Name, TriggerType, Type string
    Effect effect.Effect
}

func New(name string) Special {
    if name == "" {
        return &special{
            Effect: effect.New(""),
        }
    }
    specials := map[string]special {
        "night sky": special{
            Name: "night sky",
            BaseCooldown: 3,
            CurrentCooldown: 3,
            Effect: effect.New("night sky"),
        },
    }
    return &special{
        Name: name,
        BaseCooldown: specials[name].BaseCooldown,
        CurrentCooldown: specials[name].CurrentCooldown,
        Effect: specials[name].Effect,
    }
}

func (s *special) Trigger(base int) int {
    s.SetCurrentCooldown(s.GetBaseCooldown())
    return s.GetEffect().Proc(base)
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

func (s *special) GetEffect() effect.Effect {
    return s.Effect
}

func (s *special) GetName() string {
    return s.Name
}
