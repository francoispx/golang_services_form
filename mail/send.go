package mail

import (
	"gopkg.in/gomail.v2"
	"time"
	"contactform/config"
)

func Send(date time.Time, pdffile string, to_addr string) {
	cfg := config.GetMailConfig("config/")
	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {cfg.From},
		"To":      {to_addr},
		"Subject": {"Form PDF"},
	})
	//TODO: change body to html, change subject
	m.SetBody("text/html", "<b>Hello,</b><br>This email contains a PDF file with data from the registration form. <br> Please see the attached file.")
	m.Attach(pdffile)
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	if err := gomail.Send(s, m); err != nil {
        panic("Could not send email " + err.Error())
    }
    m.Reset()
}

