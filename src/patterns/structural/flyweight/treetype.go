package flyweight


type TreeType struct {
	name string
	color string
	texture string
}


func NewTreeType(name, color, texture string) *TreeType {
	return &TreeType{
		name: name,
		color: color,
		texture: texture,
	}
}

func (self *TreeType) Draw(canvas *Canvas, x, y float32) {
	canvas.Draw(x, y)
}

func (self *TreeType) GetName() string {
	return self.name
}

func (self *TreeType) SetName(name string) {
	self.name = name
}

func (self *TreeType) GetColor() string {
	return self.color
}

func (self *TreeType) SetColor(color string) {
	self.color = color
}

func (self *TreeType) GetTexture() string {
	return self.texture
}

func (self *TreeType) SetTexture(texture string) {
	self.texture = texture
}