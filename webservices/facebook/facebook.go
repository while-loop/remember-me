package facebook

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf/browser"
	"github.com/while-loop/remember-me/util"
	"github.com/while-loop/remember-me/webservices"
	"gopkg.in/headzoo/surf.v1"
	"net/url"
	"strings"
)

const (
	httpscheme       = "https"
	facebookhost     = "facebook.com"
	facebooklogin    = "/login.php"
	facebookchpasswd = "/settings/security/password/"
)

func init() {
	webservices.Register(facebookhost, NewFacebookWebservice())
}

type FacebookWebservice struct {
}

func NewFacebookWebservice() *FacebookWebservice {
	return &FacebookWebservice{}
}

func (f *FacebookWebservice) login(browsr *browser.Browser, email, password string) error {
	fUrl := buildFBUrl(facebooklogin)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservices.ConnectError{Hostname: fUrl.String()}
	}

	// Log in to the site.
	fm, err := browsr.Form("form[id=login_form]")
	if err != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": login_form"}
	}

	err = fm.Input("email", email)
	if err != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": login_form_email"}
	}

	err = fm.Input("pass", password)
	if err != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": login_form_password"}
	}

	if fm.Submit() != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": login_form_button"}
	}

	if strings.Contains(browsr.Title(), "Log into Facebook") {
		return webservices.AccountError{Email: email, Hostname: facebookhost}
	}

	return nil
}

func (f *FacebookWebservice) logout(browsr *browser.Browser) error {
	fUrl := buildFBUrl(facebooklogin)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservices.ConnectError{Hostname: fUrl.String()}
	}

	retErr := &webservices.AccountError{Hostname: facebookhost}
	browsr.Find("a").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		if href, exist := s.Attr("href"); exist {
			if strings.Contains(href, "logout") {
				browsr.Open(buildFBUrl(href).String())
				if err == nil {
					retErr = nil
					return false
				}
			}
		}
		return true
	})

	return retErr
}

func (f *FacebookWebservice) changePassword(browsr *browser.Browser, email, oldpasswd, newpasswd string) error {
	fUrl := buildFBUrl(facebookchpasswd)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservices.ConnectError{Hostname: fUrl.String()}
	}

	fm, err := browsr.Form("form[method=post]")
	if err != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": chpwd_form"}
	}

	err = fm.Input("password_old", oldpasswd)
	if err != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": chpw_form_old"}
	}

	err = fm.Input("password_new", newpasswd)
	if err != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": chpw_form_new"}
	}

	err = fm.Input("password_confirm", newpasswd)
	if err != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": chpw_form_conf"}
	}

	if fm.Submit() != nil {
		return webservices.ParseError{Hostname: fUrl.String() + ": login_form_button"}
	}

	// TODO put in function
	body := browsr.Body()
	err = nil
	if strings.Contains(body, "password was incorrect") {
		err = webservices.AccountError{
			Email:    email,
			Hostname: facebookhost,
		}
	} else if strings.Contains(body, "Password must differ from old password") {
		err = webservices.ChangeError{
			Hostname: facebookhost,
			Email:    email,
			Message:  "Password must differ from old password",
		}
	}
	return err
}

func (f *FacebookWebservice) ChangePassword(email, oldpasswd, newpasswd string) error {
	browsr := surf.NewBrowser()
	se := util.StickyError{}

	se.Swallow(f.login(browsr, email, oldpasswd))
	se.Swallow(f.changePassword(browsr, email, oldpasswd, newpasswd))
	se.Swallow(f.logout(browsr))

	return se.Error()
}

func buildFBUrl(path string) *url.URL {
	return &url.URL{
		Scheme: httpscheme,
		Host:   "m." + facebookhost,
		Path:   path,
	}
}
