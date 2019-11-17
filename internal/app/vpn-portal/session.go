package app

type session struct {
	IssuedOn    string
	User        string
	Profile     string
	Duration    string
	ExpiresOn   string
	ClientIP    string
	Hostname    string
	Port        string
	IssuingCA   string
	Certificate string
	PrivateKey  string
}

type sessions struct {
	Items []session
}

func (s *sessions) AddItem(mySession session) []session {
	s.Items = append(s.Items, mySession)
	return s.Items
}
