package adapter

import (
	"encoding/json"
)


type ClassAdapter struct {
	*Government
}


func NewClassAdapter() *ClassAdapter {
	return &ClassAdapter{
		&Government{},
	}
}

func (self *ClassAdapter) SendGainsInfo(data []byte) bool {
	var jsonValues map[string]any
	json.Unmarshal(data, &jsonValues)
	return self.SendGainsInformation(Tuple{TaxPayerId: jsonValues["tax_payer_id"], TaxPayerGains: jsonValues["tax_payer_gains"]})
}