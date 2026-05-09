package phonenotifer

import "log"

type PhoneService struct{}

// PhoneClient logging that client was numbered
func (ps *PhoneService) PhoneClient() error {
	log.Println("Client got numbered")
	return nil
}

func NewPhoneService() *PhoneService {
	return &PhoneService{}
}
