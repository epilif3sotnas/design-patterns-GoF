package adapter


// std
import(
	"fmt"
)


type Government struct {}


func NewGovernment() *Government {
	return &Government{}
}

func (self *Government) SendGainsInformation(data Tuple) bool {
	fmt.Println("Data successfully received")
	fmt.Println("Tax Payer ID: ", data.TaxPayerId ,"\nTax Payer Gains: ", data.TaxPayerGains)

    return true
}