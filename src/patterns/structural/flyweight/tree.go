package flyweight


type Tree struct {
	x float32
	y float32
	treeType *TreeType
}


func NewTree(x, y float32, treeType *TreeType) *Tree {
	return &Tree{
		x: x,
		y: y,
		treeType: treeType,
	}
}

func (self *Tree) Draw(canvas *Canvas) {
	self.treeType.Draw(
		canvas,
		self.x,
		self.y,
	)
}

func (self *Tree) GetX() float32 {
	return self.x
}

func (self *Tree) SetX(x float32) {
	self.x = x
}

func (self *Tree) GetY() float32 {
	return self.y
}

func (self *Tree) SetY(y float32) {
	self.y = y
}

func (self *Tree) GetTreeType() *TreeType {
	return self.treeType
}

func (self *Tree) SetTreeType(treeType *TreeType) {
	self.treeType = treeType
}