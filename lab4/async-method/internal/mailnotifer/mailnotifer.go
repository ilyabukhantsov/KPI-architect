package mailnotifer

import "log"

type Mail struct{}

// MailSomeone logging the text, that we puted to argument.
func (ml *Mail) MailSomeone(cool string) error {
	log.Println(cool)
	return nil
}

func NewMailService() *Mail {
	return &Mail{}
}
