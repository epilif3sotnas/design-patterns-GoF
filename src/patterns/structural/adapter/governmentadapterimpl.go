package adapter

import (
	"encoding/json"
)


type GovernmentAdapterImpl struct {
	government *Government
}


func NewGovernmentAdapterImpl(government *Government) *GovernmentAdapterImpl {
	return &GovernmentAdapterImpl{
		government: government,
	}
}

func (self *GovernmentAdapterImpl) SendGainsInfo(data []byte) bool {
	var jsonValues map[string]interface{}
	json.Unmarshal(data, &jsonValues)
	return self.government.SendGainsInformation(Tuple{TaxPayerId: jsonValues["tax_payer_id"], TaxPayerGains: jsonValues["tax_payer_gains"]})
}