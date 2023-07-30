package builder


type Admin struct {
	builder Builder
}


func NewAdmin(builder Builder) *Admin {
	return &Admin{
		builder: builder,
	}
}

func (self *Admin) Build() *Car {
	return self.builder.BuildCar()
}