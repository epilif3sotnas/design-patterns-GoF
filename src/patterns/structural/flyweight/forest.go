package flyweight


type Forest struct {
	trees []Tree
}


func NewForest() *Forest {
	return &Forest{
		trees: []Tree{},
	}
}

func (self *Forest) PlantTree(x, y float32, name, color, texture string) *Tree {
	var treeType *TreeType = GetInstance().GetTreeType(name, color, texture)

	tree := NewTree(x, y, treeType)

	self.trees = append(self.trees, *tree)

	return tree
}

func (self *Forest) Draw(canvas *Canvas) {
	for _, tree := range self.trees {
		tree.Draw(canvas)
	}
}

func (self *Forest) GetTrees() []Tree {
	return self.trees
}

func (self *Forest) SetTrees(trees []Tree) {
	self.trees = trees
}