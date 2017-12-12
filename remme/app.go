package remme

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"math/rand"

	"github.com/while-loop/remember-me/remme/api/services/v1/record"
	"github.com/while-loop/remember-me/remme/manager"
	"github.com/while-loop/remember-me/remme/storage"
	"github.com/while-loop/remember-me/remme/storage/stub"
	"github.com/while-loop/remember-me/remme/util"
	"github.com/while-loop/remember-me/remme/webservice"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type App struct {
	services  map[string]webservice.Webservice
	Datastore storage.DataStore
	jobsChan  chan jobRequest
	closeChan chan bool
}

const (
	workers = 10
)

var (
	emptyHost       = fmt.Errorf("empty host")
	hostDNE         = fmt.Errorf("host DNE")
	proxyParseError = fmt.Errorf("unable to change password")
)

type jobRequest struct {
	JobID      uint64
	Manager    manager.Manager
	PasswdFunc util.PasswdFunc
}

func NewApp(datastore storage.DataStore, services map[string]webservice.Webservice) *App {
	if datastore == nil {
		datastore = stub.New()
	}

	app := &App{
		Datastore: datastore,
		services:  services,
		jobsChan:  make(chan jobRequest),
		closeChan: make(chan bool),
	}

	go app.run()
	return app
}

func (a *App) run() {
	for i := 0; i < workers; i++ {
		go a.worker()
	}

	<-a.closeChan
	log.Println("app close chan recvd")

	close(a.jobsChan)
	close(a.closeChan)
}

func (a *App) ChangePasswords(mngr manager.Manager, passwdFunc util.PasswdFunc) uint64 {
	jId := rand.Uint64()
	a.jobsChan <- jobRequest{JobID: jId, Manager: mngr, PasswdFunc: passwdFunc}
	return jId
}

func (a *App) worker() {
	for j := range a.jobsChan {
		a.chPasswds(j)
	}
}

func (a *App) chPasswds(job jobRequest) {
	sites := job.Manager.GetSites()
	// mutex when updating log record
	wg := sync.WaitGroup{}
	lr, err := a.Datastore.AddLog(&storage.LogRecord{
		Time:       time.Now(),
		JobID:      job.JobID,
		User:       job.Manager.GetEmail(),
		TotalSites: uint64(len(sites)),
	})
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}

	if lr, err := a.Datastore.AddEvent(record.JobEvent{
		JobId:     lr.JobID,
		TaskId:    0,
		Type:      record.JobEvent_JOB_START,
		Email:     job.Manager.GetEmail(),
		Hostname:  "",
		Timestamp: uint64(time.Now().Unix()),
		Msg:       fmt.Sprintf("Total sites: %d", lr.TotalSites),
	}); err != nil {
		log.Printf("failed to send log record %v\n", lr)
	}

	taskId := uint64(0)
	for _, site := range sites {
		if empty([]string{site.Hostname, site.Email, site.Password}) {
			continue
		}

		service, err := a.searchService(site.Hostname)
		taskId++
		if lr, err := a.Datastore.AddEvent(record.JobEvent{
			JobId:     lr.JobID,
			TaskId:    taskId,
			Type:      record.JobEvent_TASK_START,
			Email:     site.Email,
			Hostname:  site.Hostname,
			Timestamp: uint64(time.Now().Unix()),
		}); err != nil {
			log.Printf("failed to send log record %v\n", lr)
		}

		if err != nil {
			if lr, err := a.Datastore.AddEvent(record.JobEvent{
				JobId:     lr.JobID,
				TaskId:    taskId,
				Type:      record.JobEvent_TASK_ERROR,
				Email:     site.Email,
				Hostname:  site.Hostname,
				Timestamp: uint64(time.Now().Unix()),
				Msg:       "Unsupported website",
				Version:   Version,
			}); err != nil {
				log.Printf("failed to send log record %v\n", lr)
			}
			continue
		}

		wg.Add(1)
		go a.chPasswd(job, &wg, service, site, taskId)
	}

	wg.Wait()
	lr, err = a.Datastore.UpdateLog(lr)
	fmt.Println(lr.Tries(), "/", len(lr.Failures()), "/", lr.TotalSites)
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}
}

func (a *App) chPasswd(job jobRequest, wg *sync.WaitGroup, webservice webservice.Webservice, site manager.Site, taskId uint64) {
	log.Println("Changing password for:", site.Hostname, site.Email)
	newpasswd := site.Password //job.passwdFunc() TODO
	defer wg.Done()

	event := record.JobEvent{
		JobId:    job.JobID,
		TaskId:   taskId,
		Type:     record.JobEvent_TASK_ERROR,
		Email:    site.Email,
		Hostname: site.Hostname,
		Msg:      "",
		Version:  Version,
	}

	err := webservice.ChangePassword(site.Email, site.Password, newpasswd)
	if err != nil {
		log.Println(err)
		event.Msg = err.Error()
		event.Timestamp = uint64(time.Now().Unix())
		if _, err := a.Datastore.AddEvent(event); err != nil {
			log.Printf("failed to send log record %v\n", event)
		}
		return
	}

	err = job.Manager.SavePassword(site.Hostname, site.Email, newpasswd)
	if err != nil {
		log.Printf("Failed to save password for %s %s.. reverting: %s\n", site.Hostname, site.Email, err)
		err = webservice.ChangePassword(site.Email, newpasswd, site.Password)
		if err != nil {
			// oh shit boi. failed to revert password
			log.Printf("Failed to revert back to old password for %s %s.. %s\n", site.Hostname, site.Email, err)
			event.Msg = "Failed to revert back to old password"
			event.Timestamp = uint64(time.Now().Unix())
			if _, err := a.Datastore.AddEvent(event); err != nil {
				log.Printf("failed to send log record %v\n", event)
			}
			// TODO send email to customer with new password?
			return
		}
	}

	log.Println("Password changed for:", site.Hostname, site.Email)
	event.Msg = ""
	event.Timestamp = uint64(time.Now().Unix())
	event.Type = record.JobEvent_TASK_FINISH
	if _, err := a.Datastore.AddEvent(event); err != nil {
		log.Printf("failed to send log record %v\n", event)
	}
}

func (a *App) searchService(hostname string) (webservice.Webservice, error) {
	hostname = strings.ToLower(hostname)
	hostname = strings.TrimSpace(hostname)
	if hostname == "" {
		return nil, emptyHost
	}
	for key, val := range a.services {
		if strings.Contains(key, hostname) || strings.Contains(hostname, key) {
			return val, nil
		}
	}

	return nil, hostDNE
}

func empty(str []string) bool {
	for _, s := range str {
		if strings.TrimSpace(s) == "" {
			return true
		}
	}
	return false
}
