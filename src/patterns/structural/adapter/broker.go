package adapter


type Broker struct {
	adapter GovernmentAdapter
}


func NewBroker(adapter GovernmentAdapter) *Broker {
	return &Broker{
		adapter: adapter,
	}
}

func (self *Broker) SendGains(data []byte) bool {
	return self.adapter.SendGainsInfo(data)
}