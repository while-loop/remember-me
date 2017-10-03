package webservices

var (
	defServices = map[string]Webservice{}
)

type Webservice interface {
	ChangePassword(email, oldpasswd, newpasswd string) error
}

func Register(hostname string, ws Webservice) {
	defServices[hostname] = ws
}

func Services() map[string]Webservice {
	return defServices
}
