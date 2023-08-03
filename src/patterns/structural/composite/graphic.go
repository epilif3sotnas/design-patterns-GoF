package composite


type Graphic interface {
	Move(float32, float32) bool
	Draw() bool
}