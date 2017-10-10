package remme

import (
	"fmt"
	"github.com/while-loop/remember-me/api/services/v1/changer"
	"github.com/while-loop/remember-me/storage"
	"github.com/while-loop/remember-me/manager"
	"github.com/while-loop/remember-me/webservice"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
	"github.com/while-loop/remember-me/storage/stub"
)

type App struct {
	services  map[string]webservice.Webservice
	Datastore storage.DataStore
}

var (
	emptyHost       = fmt.Errorf("empty host")
	hostDNE         = fmt.Errorf("host DNE")
	proxyParseError = fmt.Errorf("unable to change password")
)

func NewApp(datastore storage.DataStore, services map[string]webservice.Webservice) *App {
	if datastore == nil {
		datastore = &stub.StubDB{}
	}

	return &App{
		Datastore: datastore,
		services:  services,
	}
}

// TODO out = websocket/file/stdout/etc
// TODO chan out
// Status interface
// Start status
// Job start status (with subJob ID)
// Job Error status
// Job finish status
// Finish status
func (a *App) ChangePasswords(out chan<- changer.Status, mngr manager.Manager, passwdFunc PasswdFunc) {
	jId := rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
	defer func() {
		out <- newStatus(jId, 0, changer.Status_JOB_FINISH, mngr.GetEmail(), "", "")
		close(out)
	}()

	sites := mngr.GetSites()

	// mutex when updating log record
	wg := sync.WaitGroup{}
	lr, err := a.Datastore.AddLog(&storage.LogRecord{
		Time:       time.Now(),
		JobID:      jId,
		User:       mngr.GetEmail(),
		TotalSites: uint64(len(sites)),
	})
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}
	out <- newStatus(lr.JobID, 0, changer.Status_JOB_START, mngr.GetEmail(), "",
		fmt.Sprintf("Total sites: %d", lr.TotalSites))

	taskId := uint64(0)
	for _, site := range sites {
		if empty([]string{site.Hostname, site.Email, site.Password}) {
			continue
		}

		service, err := a.searchService(site.Hostname)

		taskId++
		out <- newStatus(lr.JobID, taskId, changer.Status_TASK_START, site.Email, site.Hostname, "")

		if err != nil {
			lr.AddFailure(site.Hostname, site.Email, "Unsupported website", Version)
			out <- newStatus(lr.JobID, taskId, changer.Status_TASK_ERROR, site.Email, site.Hostname, "Unsupported website")
			continue
		}

		wg.Add(1)
		go chPasswd(out, &wg, service, mngr, site, lr, taskId)
	}

	wg.Wait()
	lr, err = a.Datastore.UpdateLog(lr)
	fmt.Println(lr.Tries(), "/", len(lr.Failures()), "/", lr.TotalSites)
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}
}

func newStatus(jId, tId uint64, typ changer.Status_Type, email, hname, msg string) changer.Status {
	return changer.Status{JobId: jId, TaskId: tId, Type: typ, Email: email, Hostname: hname, Msg: msg}
}

func chPasswd(out chan<- changer.Status, wg *sync.WaitGroup, goservice webservice.Webservice,
	mngr manager.Manager, gosite manager.Site, lr *storage.LogRecord, goTaskId uint64) {

	lr.IncTries(1)
	log.Println("Changing password for:", gosite.Hostname, gosite.Email)
	newpasswd := gosite.Password //passwdFunc()) TODO
	defer wg.Done()

	err := goservice.ChangePassword(gosite.Email, gosite.Password, newpasswd)
	if err != nil {
		log.Println(err)
		lr.AddFailure(gosite.Hostname, gosite.Email, err.Error(), Version)
		if _, ok := err.(webservice.ParseError); ok {
			err = proxyParseError // user-friendly error
		}

		out <- newStatus(lr.JobID, goTaskId, changer.Status_TASK_ERROR, gosite.Email, gosite.Hostname, err.Error())
		return
	}

	err = mngr.SavePassword(gosite.Hostname, gosite.Email, newpasswd)
	if err != nil {
		log.Printf("Failed to save password for %s %s.. reverting: %s\n", gosite.Hostname, gosite.Email, err)
		lr.AddFailure(gosite.Hostname, gosite.Email, "Failed to save new password", Version)
		err = goservice.ChangePassword(gosite.Email, newpasswd, gosite.Password)
		if err != nil {
			// oh shit boi. failed to revert password
			log.Printf("Failed to revert back to old password for %s %s.. %s\n", gosite.Hostname, gosite.Email, err)
			lr.AddFailure(gosite.Hostname, gosite.Email, "Failed to revert back to old password", Version)
			if _, ok := err.(webservice.ParseError); ok {
				err = proxyParseError // user-friendly error
			}

			// TODO send email to customer with new password?
			out <- newStatus(lr.JobID, goTaskId, changer.Status_TASK_ERROR, gosite.Email, gosite.Hostname, "Failed to revert back to old password")
			return
		}
	}

	log.Println("Password changed for:", gosite.Hostname, gosite.Email)
	out <- newStatus(lr.JobID, goTaskId, changer.Status_TASK_FINISH, gosite.Email, gosite.Hostname, "")
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
