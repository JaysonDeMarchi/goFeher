package effect

type Effect interface {
    Proc(int) int
    GetOperator() string
    GetStat() string
    GetTrigger() string
    GetType() string
    GetValue() float64
}
