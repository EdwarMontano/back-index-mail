package enronmail

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

type Item struct {
	IdMsg      string `json:"IdMsg"`
	DateMsg    string `json:"DateMsg"`
	FromMsg    string `json:"FromMsg"`
	ToMsg      string `json:"ToMsg"`
	SubjectMsg string `json:"SubjectMsg"`
	CcMsg      string `json:"CcMsg"`
	BccMsg     string `json:"BccMsg"`
	XFromMsg   string `json:"XFromMsg"`
	XToMsg     string `json:"XToMsg"`
	XccMsg     string `json:"XccMsg"`
	XbccMsg    string `json:"XbccMsg"`
	Content    string `json:"ContentMsg"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
