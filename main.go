package main

import (
	"./manager"
	"./webservice"
	"fmt"
	"log"
)

func main() {

	memMan := genData()
	fb := webservice.NewFacebookWebservice()
	services := map[string]webservice.Webservice{
		fb.GetDomain(): fb,
	}

	for _, site := range memMan.GetSites() {
		service := services[site.Domain]
		if service == nil {
			continue
		}

		fmt.Println("Changing password for:", site.Domain, site.Email)

		err := service.Login(site.Email, site.Password)
		if err != nil {
			log.Println(err)
			continue
		}

		err = service.ChangePassword(site.Email, site.Password, genPasswd())
		if err != nil {
			log.Println(err)
			continue
		}

		err = service.Logout()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("Password changed for:", site.Domain, site.Email)

	}
}
func genPasswd() string {
	return "newpassword"
}
func genData() manager.Manager {
	mm := manager.NewMemManager()
	mm.SavePassword("facebook.com", "email", "password")
	return mm
}
