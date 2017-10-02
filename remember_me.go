package remember

import (
	"./db"
	"./manager"
	"./webservice"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	Version = "0.0.1"
)

type PasswdFunc func() string

var DefaultPasswdFunc = func() string {
	return NewPasswordGen(32, true, true).Generate()
}

type App struct {
	services  map[string]webservice.Webservice
	datastore db.DataStore
}

func NewApp(datastore db.DataStore, services ...webservice.Webservice) *App {
	a := &App{
		datastore: datastore,
		services:  map[string]webservice.Webservice{},
	}
	a.AddServices(services...)
	return a
}

func (a *App) AddServices(services ...webservice.Webservice) {
	for _, w := range services {
		a.services[strings.ToLower(w.GetHostname())] = w
	}
}

// TODO out = websocket
func (a *App) ChangePasswords(out io.Writer, mngr manager.Manager, passwdFunc PasswdFunc) {
	sites := mngr.GetSites()

	// mutex when updating log record
	wg := sync.WaitGroup{}
	lr, err := a.datastore.AddLog(&db.LogRecord{
		Time:       time.Now(),
		JobID:      rand.New(rand.NewSource(time.Now().UnixNano())).Uint64(),
		User:       mngr.GetEmail(),
		TotalSites: uint(len(sites)),
	})

	if err != nil {
		log.Println("Unable to save log", err, lr)
	}

	for _, site := range sites {
		service := a.searchService(site.Hostname)
		if service == nil {
			lr.AddFailure(site.Hostname, site.Email, "Unsupported host", Version)
			continue
		}

		wg.Add(1)
		go func(goservice webservice.Webservice, gosite manager.Site) {
			lr.IncTries(1)
			log.Println("Changing password for:", gosite.Hostname, gosite.Email)

			err := goservice.ChangePassword(gosite.Email, gosite.Password, gosite.Password) //passwdFunc()) TODO
			if err != nil {
				log.Println(err)
				lr.AddFailure(gosite.Hostname, gosite.Email, err.Error(), Version)
				wg.Done()
				return
			}
			log.Println("Password changed for:", gosite.Hostname, gosite.Email)
		}(service, site)
	}

	wg.Wait()
	lr, err = a.datastore.UpdateLog(lr)
	fmt.Println(lr.Tries(), "/", len(lr.Failures()), "/", lr.TotalSites)
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}
}

func (a *App) searchService(hostname string) webservice.Webservice {
	hostname = strings.ToLower(hostname)
	hostname = strings.TrimSpace(hostname)
	if hostname == "" {
		return nil
	}
	for key, val := range a.services {
		if strings.Contains(key, hostname) || strings.Contains(hostname, key) {
			return val
		}
	}

	return nil
}
