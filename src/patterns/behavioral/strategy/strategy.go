package strategy


type Strategy interface {
	Execute(a int32, b int32) float64
}