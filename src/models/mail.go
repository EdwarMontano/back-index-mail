package models

type Getter interface {
	GetAll() []Mail
}

type Adder interface {
	Add(item Mail)
}

type Mail struct {
	Bcc     string `json:"Bcc"`
	Cc      string `json:"Cc"`
	Content string `json:"Content"`
	DateMSG string `json:"DateMSG"`
	From    string `json:"From"`
	IDMSG   string `json:"IdMSG"`
	Path    string `json:"Path"`
	Subject string `json:"Subject"`
	To      string `json:"To"`
	XFrom   string `json:"X-From"`
	XTo     string `json:"X-To"`
	XBcc    string `json:"X-bcc"`
	XCc     string `json:"X-cc"`
	// CreatedAt time.Time `json:created_at`
}

type Repo struct {
	Mails []Mail
}

func New() *Repo {
	return &Repo{
		Mails: []Mail{},
	}
}

func (r *Repo) Add(mail Mail) {
	r.Mails = append(r.Mails, mail)
}

func (r *Repo) GetAll() []Mail {
	return r.Mails
}
