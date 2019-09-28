package effect

type effect struct {
    Operator string
    Stat string
    Trigger string
    Type string
    Value float64
}

func New(name string) Effect {
    if (name == "") {
        return &effect{}
    }
    effects := map[string]effect {
        "night sky": effect{
            Operator: "mult",
            Stat: "damage",
            Trigger: "attack",
            Type: "buff",
            Value: 0.5,
        },
    }

    return &effect{
        Operator: effects[name].Operator,
        Stat: effects[name].Stat,
        Trigger: effects[name].Trigger,
        Type: effects[name].Type,
        Value: effects[name].Value,
    }
}

func (e *effect) Proc(base int) int {
    result := 0
    if e.GetType() == "buff" {
        boost := e.applyBoost(base)
        result = base + boost
    }
    return result
}

func (e *effect) applyBoost(base int) int {
    result := 0.0
    if e.GetOperator() == "mult" {
        result = float64(base) * e.GetValue()
    }
    return int(result)
}

func (e *effect) GetOperator() string {
    return e.Operator
}

func (e *effect) GetStat() string {
    return e.Stat
}

func (e *effect) GetTrigger() string {
    return e.Trigger
}

func (e *effect) GetType() string {
    return e.Type
}

func (e *effect) GetValue() float64 {
    return e.Value
}
