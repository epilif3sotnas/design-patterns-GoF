package prototype

// std
import(
	"strconv"
)


type Clothing struct {
	brand string
	model string
	price float32
	owner string
}


func NewClothing(brand string, model string, price float32, owner string) *Clothing {
	return &Clothing{
		brand: brand,
		model: model,
		price: price,
		owner: owner,
	}
}

func (self *Clothing) GetClothingInfo() string {
	return "Brand: " + self.brand + "\nModel: " + self.model + "\nPrice: " + strconv.FormatFloat(float64(self.price), 'f', -1, 32) + "\nOwner: " + self.owner
}

func (self *Clothing) Copy() *Clothing {
	panic("Method `Copy` not implemented")
}

func (self *Clothing) GetBrand() string {
	return self.brand
}

func (self *Clothing) SetBrand(brand string) {
	self.brand = brand
}

func (self *Clothing) GetModel() string {
	return self.model
}

func (self *Clothing) SetModel(model string) {
	self.model = model
}

func (self *Clothing) GetPrice() float32 {
	return self.price
}

func (self *Clothing) SetPrice(price float32) {
	self.price = price
}

func (self *Clothing) GetOwner() string {
	return self.owner
}

func (self *Clothing) SetOwner(owner string) {
	self.owner = owner
}