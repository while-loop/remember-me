package remme

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"math/rand"

	"github.com/while-loop/remember-me/remme/api/services/v1/record"
	"github.com/while-loop/remember-me/remme/log"
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
	log.Info("app close chan recvd")

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
	event := record.JobEvent{
		JobId:     job.JobID,
		TaskId:    0,
		Type:      record.JobEvent_JOB_START,
		Email:     job.Manager.GetEmail(),
		Hostname:  job.Manager.Name(),
		Timestamp: uint64(time.Now().Unix()),
		Version:   Version,
	}

	sites, err := job.Manager.GetSites()
	if err != nil {
		log.Error("chPasswds: getSites", err)
		event.Type = record.JobEvent_JOB_FINISH
		event.Msg = err.Error()
		if _, err := a.Datastore.AddEvent(event); err != nil {
			log.Errorf("failed to send log record %v\n", event)
		}
		return
	}
	// mutex when updating log record
	wg := sync.WaitGroup{}
	event.Msg = fmt.Sprintf("Total sites: %d", len(sites))
	if _, err := a.Datastore.AddEvent(event); err != nil {
		log.Errorf("failed to send log record %v\n", event)
	}

	taskId := uint64(0)
	for _, site := range sites {
		if empty([]string{site.Hostname, site.Email, site.Password}) {
			continue
		}

		service, err := a.searchService(site.Hostname)
		taskId++
		event.Email = site.Email
		event.Hostname = site.Hostname
		event.Timestamp = uint64(time.Now().Unix())
		event.TaskId = taskId
		if err != nil {
			event.Type = record.JobEvent_TASK_ERROR
			event.Msg = "Unsupported website"
			if _, err := a.Datastore.AddEvent(event); err != nil {
				log.Errorf("failed to send log record %v\n", event)
			}
			continue
		}
		event.Type = record.JobEvent_TASK_START
		if _, err := a.Datastore.AddEvent(event); err != nil {
			log.Errorf("failed to send job event %v\n", event)
		}

		wg.Add(1)
		go a.chPasswd(job, &wg, service, site, taskId)
	}

	wg.Wait()
	event.TaskId = 0
	event.Type = record.JobEvent_JOB_FINISH
	event.Email = job.Manager.GetEmail()
	event.Hostname = job.Manager.Name()
	event.Timestamp = uint64(time.Now().Unix())
	event.Msg = fmt.Sprintf("Total sites: %d", len(sites))
	if _, err := a.Datastore.AddEvent(event); err != nil {
		log.Errorf("failed to send log record %v\n", event)
	}
	log.Infof("Job finished %v\n", event)
}

func (a *App) chPasswd(job jobRequest, wg *sync.WaitGroup, webservice webservice.Webservice, site manager.Site, taskId uint64) {
	log.Info("Changing password for:", site.Hostname, site.Email)
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
		log.Error(err)
		event.Msg = err.Error()
		event.Timestamp = uint64(time.Now().Unix())
		if _, err := a.Datastore.AddEvent(event); err != nil {
			log.Errorf("failed to send log record %v\n", event)
		}
		return
	}

	err = job.Manager.SavePassword(site.Hostname, site.Email, newpasswd)
	if err != nil {
		log.Errorf("Failed to save password for %s %s.. reverting: %s\n", site.Hostname, site.Email, err)
		err = webservice.ChangePassword(site.Email, newpasswd, site.Password)
		if err != nil {
			// oh shit boi. failed to revert password
			log.Errorf("Failed to revert back to old password for %s %s.. %s\n", site.Hostname, site.Email, err)
			event.Msg = "Failed to revert back to old password"
			event.Timestamp = uint64(time.Now().Unix())
			if _, err := a.Datastore.AddEvent(event); err != nil {
				log.Errorf("failed to send log record %v\n", event)
			}
			// TODO send email to customer with new password?
			return
		}
	}

	log.Info("Password changed for:", site.Hostname, site.Email)
	event.Msg = ""
	event.Timestamp = uint64(time.Now().Unix())
	event.Type = record.JobEvent_TASK_FINISH
	if _, err := a.Datastore.AddEvent(event); err != nil {
		log.Errorf("failed to send log record %v\n", event)
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
