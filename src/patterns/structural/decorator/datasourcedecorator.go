package decorator


type DataSourceDecorator struct {
	wrappee DataSource
}


func NewDataSourceDecorator(wrappee DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{
		wrappee: wrappee,
	}
}

func (self *DataSourceDecorator) WriteData(data []byte) {
	self.wrappee.WriteData(data)
}

func (self *DataSourceDecorator) ReadData() []byte {
	return self.wrappee.ReadData()
}

func (self *DataSourceDecorator) GetWrappee() DataSource {
	return self.wrappee
}

func (self *DataSourceDecorator) SetWrappee(wrappee DataSource) {
	self.wrappee = wrappee
}