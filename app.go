package remme

import (
	"github.com/while-loop/remember-me/db"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
	"github.com/while-loop/remember-me/webservices"
	"github.com/while-loop/remember-me/managers"
	"github.com/while-loop/remember-me/api/services/v1/changer"
)

type App struct {
	services  map[string]webservices.Webservice
	Datastore db.DataStore
}

var (
	emptyHost = fmt.Errorf("empty host")
	hostDNE   = fmt.Errorf("host DNE")
)

func NewApp(datastore db.DataStore, services map[string]webservices.Webservice) *App {
	if datastore == nil {
		datastore = &db.StubDB{}
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
func (a *App) ChangePasswords(out chan<- changer.Status, mngr managers.Manager, passwdFunc PasswdFunc) {
	sites := mngr.GetSites()

	// mutex when updating log record
	wg := sync.WaitGroup{}
	lr, err := a.Datastore.AddLog(&db.LogRecord{
		Time:       time.Now(),
		JobID:      rand.New(rand.NewSource(time.Now().UnixNano())).Uint64(),
		User:       mngr.GetEmail(),
		TotalSites: uint64(len(sites)),
	})
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}
	out <- changer.Status{
		JobId:    lr.JobID,
		Type:     changer.Status_JOB_START,
		Email:    mngr.GetEmail(),
		Hostname: "",
		Msg:      fmt.Sprintf("Total sites: %d", lr.TotalSites),
		TaskId:   uint64(0),
	}

	taskId := uint64(0)
	for _, site := range sites {
		if empty([]string{site.Hostname, site.Email, site.Password}) {
			continue
		}

		service, err := a.searchService(site.Hostname)

		taskId++
		out <- changer.Status{
			JobId:    lr.JobID,
			TaskId:   taskId,
			Type:     changer.Status_TASK_START,
			Email:    site.Email,
			Hostname: site.Hostname,
			Msg:      "",
		}

		if err != nil {
			lr.AddFailure(site.Hostname, site.Email, "Unsupported website", Version)
			out <- changer.Status{
				JobId:    lr.JobID,
				TaskId:   taskId,
				Type:     changer.Status_TASK_ERROR,
				Email:    site.Email,
				Hostname: site.Hostname,
				Msg:      "Unsupported website",
			}
			continue
		}

		wg.Add(1)
		go func(goservice webservices.Webservice, gosite managers.Site, goTaskId uint64) {
			lr.IncTries(1)
			log.Println("Changing password for:", gosite.Hostname, gosite.Email)
			newpasswd := gosite.Password //passwdFunc()) TODO
			defer wg.Done()

			err := goservice.ChangePassword(gosite.Email, gosite.Password, newpasswd)
			if err != nil {
				log.Println(err)
				lr.AddFailure(gosite.Hostname, gosite.Email, err.Error(), Version)
				out <- changer.Status{
					JobId:    lr.JobID,
					TaskId:   goTaskId,
					Type:     changer.Status_TASK_ERROR,
					Email:    gosite.Email,
					Hostname: gosite.Hostname,
					Msg:      err.Error(),
				}
				return
			}

			err = mngr.SavePassword(gosite.Hostname, gosite.Email, newpasswd)
			if err != nil {
				log.Printf("Failed to save password for %s %s.. reverting: %s\n", gosite.Hostname, gosite.Email, err)
				lr.AddFailure(gosite.Hostname, gosite.Email, "Failed to save new password", Version)
				err = goservice.ChangePassword(gosite.Email, newpasswd, gosite.Password)
				if err != nil {
					// oh shit boi.
					log.Println(err)
					lr.AddFailure(gosite.Hostname, gosite.Email, "Failed to revert back to old password", Version)
					log.Printf("Failed to revert back to old password for %s %s.. %s\n", gosite.Hostname, gosite.Email, err)
					// TODO send email to customer with new password?
					out <- changer.Status{
						JobId:    lr.JobID,
						TaskId:   goTaskId,
						Type:     changer.Status_TASK_ERROR,
						Email:    gosite.Email,
						Hostname: gosite.Hostname,
						Msg:      "Failed to revert back to old password",
					}
					return
				}
			}

			log.Println("Password changed for:", gosite.Hostname, gosite.Email)
			out <- changer.Status{
				JobId:    lr.JobID,
				TaskId:   goTaskId,
				Type:     changer.Status_TASK_FINISH,
				Email:    gosite.Email,
				Hostname: gosite.Hostname,
				Msg:      "",
			}
		}(service, site, taskId)
	}

	wg.Wait()
	lr, err = a.Datastore.UpdateLog(lr)
	fmt.Println(lr.Tries(), "/", len(lr.Failures()), "/", lr.TotalSites)
	if err != nil {
		log.Println("Unable to save log", err, lr)
	}
	out <- changer.Status{
		JobId:    lr.JobID,
		TaskId:   0,
		Type:     changer.Status_JOB_FINISH,
		Email:    mngr.GetEmail(),
		Hostname: "",
		Msg:      "",
	}
	close(out)
}

func (a *App) searchService(hostname string) (webservices.Webservice, error) {
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
