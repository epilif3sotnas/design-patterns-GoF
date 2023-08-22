package flyweight


type TreeFactory struct {
	treeTypes []TreeType
}


var instance *TreeFactory


func newTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: []TreeType{},
	}
}

func GetInstance() *TreeFactory {
	if instance == nil {
		instance = newTreeFactory()
		return instance
	}
	return instance
}

func (self *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	var treeType *TreeType
	if len(self.treeTypes) == 0 {
		treeType = NewTreeType(name, color, texture)
		_ = append(self.treeTypes, *treeType)
		return treeType
	}

	for _, treeTypeIter := range self.treeTypes {
		if name == treeTypeIter.GetName() {
			treeType = &treeTypeIter
			break
		}
	}

	if treeType == nil {
		treeType = NewTreeType(name, color, texture)
		_ = append(self.treeTypes, *treeType)
	}
	
	return treeType
}

func (self *TreeFactory) GetTreeTypes() []TreeType {
	return self.treeTypes
}

func (self *TreeFactory) SetTreeType(treeTypes []TreeType) {
	self.treeTypes = treeTypes
}