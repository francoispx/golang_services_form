package mail

import (
	"gopkg.in/gomail.v2"
	"time"
	"bgtasks/config"
)

func Send(date time.Time, csvfile string) {
	cfg := config.GetMailConfig("../config")

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From":    {cfg.From},
		"To":      cfg.Tolist,
		"Subject": {"Sign ups " + date.Format("Mon Jan 2 2006")},
	})

	m.SetBody("text/html", "All<br><br>This email contains a CSV of new signups since " + date.Format("Mon Jan 2 2006 15:04:05") + "<br><br>See attached for details ...")
	m.Attach(csvfile)

	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	s, err := d.Dial()
	// Send the email to Bob, Cora and Dan.
	if err != nil {
		panic(err)
	}
	if err := gomail.Send(s, m); err != nil {
        panic("Could not send email " + err.Error())
    }
    m.Reset()

}

