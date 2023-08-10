package decorator


type DataSource interface {
	WriteData([]byte)
	ReadData() []byte
}