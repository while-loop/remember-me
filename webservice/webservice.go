package webservice

var (
	defServices = []Webservice{}
)

type Webservice interface {
	ChangePassword(email, oldpasswd, newpasswd string) error
	GetHostname() string
}

func Register(ws Webservice) {
	defServices = append(defServices, ws)
}

func Services() []Webservice {
	return defServices
}
