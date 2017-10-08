package lastpass

import (
	"github.com/while-loop/lastpass-go"
	"github.com/while-loop/remember-me/managers"
	"log"
	"net/url"
	"strings"
)

type LastPassManager struct {
	lp    *lastpass.LastPass
	email string
}

const (
	name = "lastpass"
)

func init() {
	managers.Register(name, func(email, password string) (managers.Manager, error) {
		return NewLastPassManager(email, password)
	})
}

func NewLastPassManager(username, password string) (*LastPassManager, error) {
	lp, err := lastpass.New(username, password)
	return &LastPassManager{
		lp:    lp,
		email: username,
	}, err
}

func (lp LastPassManager) GetEmail() string {
	return lp.email
}

func (lp *LastPassManager) GetPassword(hostname, email string) (string, error) {
	hostname = strings.ToLower(hostname)
	accs, err := lp.lp.GetAccounts()
	if err != nil {
		return "", err
	}

	for _, acc := range accs {
		// TODO regex?
		u := strings.ToLower(acc.Url)
		if strings.Contains(u, hostname) && strings.EqualFold(email, acc.Username) {
			return acc.Password, nil
		}
	}
	return "", managers.AccountDNE(hostname, email)
}

func (lp *LastPassManager) SavePassword(hostname, email, password string) error {
	return nil
}

func (lp *LastPassManager) GetSites() []managers.Site {
	sites := []managers.Site{}

	accs, err := lp.lp.GetAccounts()
	if err != nil {
		return sites
	}

	for _, acc := range accs {
		u, err := url.Parse(acc.Url)
		if err != nil {
			log.Println("Failed to parse URL: ", acc.Url, err)
		}

		sites = append(sites, managers.Site{
			Hostname: strings.ToLower(u.Hostname()),
			Email:    acc.Username,
			Password: acc.Password,
		})
	}

	return sites
}
