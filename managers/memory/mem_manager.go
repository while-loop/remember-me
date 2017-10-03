package memory

import (
	"github.com/while-loop/remember-me/managers"
)

const (
	name = "mem"
)

func init() {
	managers.Register(name, func(email, password string) (managers.Manager, error) {
		return NewMemManager(), nil
	})
}

type MemManager struct {
	// passwds[hostname][email] = password
	passwds map[string]map[string]string
}

func NewMemManager() *MemManager {
	return &MemManager{
		passwds: map[string]map[string]string{},
	}
}

func (m *MemManager) GetPassword(hostname, email string) (string, error) {
	if _, ok := m.passwds[hostname]; !ok {
		return "", managers.AccountDNE(hostname, email)
	}

	if passwd, ok := m.passwds[hostname][email]; !ok {
		return "", managers.AccountDNE(hostname, email)
	} else {
		return passwd, nil
	}
}

func (m *MemManager) GetEmail() string {
	return "mem"
}

func (m *MemManager) SavePassword(hostname, email, password string) error {
	if _, ok := m.passwds[hostname]; !ok {
		m.passwds[hostname] = map[string]string{}
	}

	m.passwds[hostname][email] = password
	return nil
}

func (m *MemManager) GetSites() []managers.Site {
	sites := []managers.Site{}
	for host, emails := range m.passwds {
		for email, passwd := range emails {
			sites = append(sites, managers.Site{
				Hostname: host,
				Email:    email,
				Password: passwd,
			})
		}
	}

	return sites
}
