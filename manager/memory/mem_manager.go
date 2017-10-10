package memory

import (
	"github.com/while-loop/remember-me/manager"
)

const (
	name = "mem"
)

func init() {
	manager.Register(name, func(email, password string) (manager.Manager, error) {
		return New(), nil
	})
}

type MemManager struct {
	// passwds[hostname][email] = password
	passwds map[string]map[string]string
}

func New() manager.Manager {
	return &MemManager{
		passwds: map[string]map[string]string{},
	}
}

func (m *MemManager) GetPassword(hostname, email string) (string, error) {
	if _, ok := m.passwds[hostname]; !ok {
		return "", manager.AccountDNE(hostname, email)
	}

	if passwd, ok := m.passwds[hostname][email]; !ok {
		return "", manager.AccountDNE(hostname, email)
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

func (m *MemManager) GetSites() []manager.Site {
	sites := []manager.Site{}
	for host, emails := range m.passwds {
		for email, passwd := range emails {
			sites = append(sites, manager.Site{
				Hostname: host,
				Email:    email,
				Password: passwd,
			})
		}
	}

	return sites
}
