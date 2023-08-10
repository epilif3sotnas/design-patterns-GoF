package decorator


import(
	"encoding/base64"
	"fmt"
)


type EncryptionDecorator struct {
	*DataSourceDecorator
}


func NewEncryptionDecorator(wrappee DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{
		&DataSourceDecorator{
			wrappee: wrappee,
		},
	}
}

func (self *EncryptionDecorator) WriteData(data []byte) {
	base64Encoded := base64.StdEncoding.EncodeToString(data)

	self.wrappee.WriteData([]byte(base64Encoded))
}

func (self *EncryptionDecorator) ReadData() []byte {
	base64Encoded := self.wrappee.ReadData()

	infoDecoded, err := base64.StdEncoding.DecodeString(string(base64Encoded[:]))

	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	return infoDecoded
}