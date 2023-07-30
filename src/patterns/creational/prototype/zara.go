package prototype


type Zara struct {
	*Clothing
}


func NewZara(brand string, model string, price float32, owner string) *Zara {
	return &Zara{
		&Clothing{
			brand: brand,
			model: model,
			price: price,
			owner: owner,
		},
	}
}

func (self *Zara) Copy() *Zara {
	return &Zara{
		&Clothing{
			brand: self.GetBrand(),
			model: self.GetModel(),
			price: self.GetPrice(),
			owner: self.GetOwner(),
		},
	}
}