package manager

import (
	"github.com/mattn/lastpass-go"
	"strings"
	"net/url"
	"log"
)

type LastPassManager struct {
	vault *lastpass.Vault
}

func NewLastPassManager(username, password string) (*LastPassManager, error) {
	v, err := lastpass.CreateVault(username, password)
	return &LastPassManager{
		vault: v,
	}, err
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
			Hostname: u.Hostname(),
			Email:    acc.Username,
			Password: acc.Password,
		})
	}

	return sites
}
