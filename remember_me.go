package main

import (
	"./manager"
	"./webservice"
	"./db"
	"log"
	"time"
	"math/rand"
	"net"
	"sync"
)

type App struct {
	services  map[string]webservice.Webservice
	datastore db.DataStore
}

func NewApp(datastore db.DataStore, services ... webservice.Webservice) *App {
	a := &App{
		datastore: datastore,
		services:  map[string]webservice.Webservice{},
	}
	a.AddServices(services...)
	return a
}

func (a *App) AddServices(services ... webservice.Webservice) {
	for _, w := range services {
		a.services[w.GetHostname()] = w
	}
}

// TODO conn = websocket
func (a *App) ChangePasswords(conn net.Conn, mngr manager.Manager) {

	sites := mngr.GetSites()

	// mutex when updating log record
	wg := sync.WaitGroup{}
	lrmu := sync.Mutex{}
	lr, err := a.datastore.AddLog(db.LogRecord{
		Time:       time.Now(),
		JobID:      rand.New(rand.NewSource(time.Now().UnixNano())).Uint64(),
		User:       "email",
		TotalSites: uint(len(sites)),
		Tries:      0,
	})

	if err != nil {
		log.Println("Unable to save log", err, lr)
	}

	for _, site := range sites {
		service := services[site.Hostname]
		if service == nil {
			lrmu.Lock()
			lr.AddFailure(site.Hostname, site.Email, "Unsupported host", VERSION)
			lrmu.Unlock()
			continue
		}
		wg.Add(1)
		go func(goservice webservice.Webservice) {
			lrmu.Lock()
			lr.Tries++
			lrmu.Unlock()
			log.Println("Changing password for:", site.Hostname, site.Email)

			err := service.Login(site.Email, site.Password)
			if err != nil {
				log.Println(err)
				lrmu.Lock()
				lr.Fails++
				lr.AddFailure(site.Hostname, site.Email, err.Error(), VERSION)
				lrmu.Unlock()
				wg.Done()
				return
			}

			err = service.ChangePassword(site.Email, site.Password, genPasswd())
			if err != nil {
				log.Println(err)
				lrmu.Lock()
				lr.Fails++
				lr.AddFailure(site.Hostname, site.Email, err.Error(), VERSION)
				lrmu.Unlock()
				wg.Done()
				return
			}

			err = service.Logout()
			if err != nil {
				log.Println(err)
				lrmu.Lock()
				lr.Fails++
				lr.AddFailure(site.Hostname, site.Email, err.Error(), VERSION)
				lrmu.Unlock()
				wg.Done()
				return
			}


			log.Println("Password changed for:", site.Hostname, site.Email)
		}(service)
	}

	wg.Wait()
	lr, err = a.datastore.UpdateLog(lr)
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}
}
