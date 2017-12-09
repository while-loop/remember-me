package lastpass

import (
	"log"
	"net/url"
	"strings"

	"github.com/while-loop/lastpass-go"
	"github.com/while-loop/remember-me/remme/manager"
)

type lastPassManager struct {
	lp    *lastpass.Vault
	email string
}

const (
	name = "lastpass"
)

func init() {
	manager.Register(name, func(email, password string) (manager.Manager, error) {
		return New(email, password)
	})
}

func New(username, password string) (manager.Manager, error) {
	lp, err := lastpass.New(username, password)
	return &lastPassManager{
		lp:    lp,
		email: username,
	}, err
}

func (lp lastPassManager) GetEmail() string {
	return lp.email
}

func (lp *lastPassManager) GetPassword(hostname, email string) (string, error) {
	hostname = strings.ToLower(hostname)
	acc, err := lp.getAccount(hostname, email)
	if err != nil {
		return "", err
	}

	return acc.Password, nil
}

func (lp *lastPassManager) getAccount(hostname, email string) (*lastpass.Account, error) {
	accs, err := lp.lp.GetAccounts()
	if err != nil {
		return nil, err
	}

	for _, acc := range accs {
		// TODO regex?
		u := strings.ToLower(acc.Url)
		if strings.Contains(u, hostname) && strings.EqualFold(email, acc.Username) {
			return acc, nil
		}
	}

	return nil, manager.AccountDNE(hostname, email)
}

func (lp *lastPassManager) SavePassword(hostname, email, password string) error {
	acc, err := lp.getAccount(hostname, email)
	if err != nil {
		return err
	}

	acc.Password = password
	_, err = lp.lp.UpdateAccount(acc)
	return err
}

func (lp *lastPassManager) GetSites() []manager.Site {
	sites := []manager.Site{}

	accs, err := lp.lp.GetAccounts()
	if err != nil {
		return sites
	}

	for _, acc := range accs {
		u, err := url.Parse(acc.Url)
		if err != nil {
			log.Println("Failed to parse URL: ", acc.Url, err)
		}

		sites = append(sites, manager.Site{
			Hostname: strings.ToLower(u.Hostname()),
			Email:    acc.Username,
			Password: acc.Password,
		})
	}

	return sites
}
