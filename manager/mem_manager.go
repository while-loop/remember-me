package manager

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
		return "", AccountDNE(hostname, email)
	}

	if passwd, ok := m.passwds[hostname][email]; !ok {
		return "", AccountDNE(hostname, email)
	} else {
		return passwd, nil
	}
}

func (m *MemManager) SavePassword(hostname, email, password string) error {
	if _, ok := m.passwds[hostname]; !ok {
		m.passwds[hostname] = map[string]string{}
	}

	m.passwds[hostname][email] = password
	return nil
}

func (m *MemManager) GetSites() []Site {
	sites := []Site{}
	for host, emails := range m.passwds {
		for email, passwd := range emails {
			sites = append(sites, Site{
				Hostname:   host,
				Email:    email,
				Password: passwd,
			})
		}
	}

	return sites
}
