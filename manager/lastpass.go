package manager

import (
	"github.com/mattn/lastpass-go"
	"log"
	"net/url"
	"strings"
)

type LastPassManager struct {
	vault *lastpass.Vault
	email string
}

const (
	name = "lastpass"
)

func init() {
	Register(name, func(email, password string) (Manager, error) {
		return NewLastPassManager(email, password)
	})
}

func NewLastPassManager(username, password string) (*LastPassManager, error) {
	v, err := lastpass.CreateVault(username, password)
	return &LastPassManager{
		vault: v,
		email: username,
	}, err
}

func (lp LastPassManager) GetEmail() string {
	return lp.email
}

func (lp *LastPassManager) GetPassword(hostname, email string) (string, error) {
	for _, acc := range lp.vault.Accounts {
		// TODO regex?
		if strings.Contains(acc.Url, hostname) && email == acc.Username {
			return acc.Password, nil
		}
	}
	return "", AccountDNE(hostname, email)
}

func (lp *LastPassManager) SavePassword(hostname, email, password string) error {
	return nil
}

func (lp *LastPassManager) GetSites() []Site {
	sites := []Site{}
	for _, acc := range lp.vault.Accounts {
		u, err := url.Parse(acc.Url)
		if err != nil {
			log.Println("Failed to parse URL: ", acc.Url, err)
		}

		sites = append(sites, Site{
			Hostname: strings.ToLower(u.Hostname()),
			Email:    acc.Username,
			Password: acc.Password,
		})
	}

	return sites
}
