package webservice

import (
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

const (
	httpscheme       = "https"
	facebookhost     = "facebook.com"
	facebooklogin    = "/login.php"
	facebookchpasswd = "/settings/security/password/"
)

type FacebookWebservice struct {
	browser *browser.Browser
}

func NewFacebookWebservice() *FacebookWebservice {
	return &FacebookWebservice{surf.NewBrowser()}
}

func (f *FacebookWebservice) Login(email, password string) error {
	fUrl := buildUrl(facebooklogin)

	err := f.browser.Open(fUrl.String())
	if err != nil {
		return ConnectError(fUrl.String())
	}

	// Log in to the site.
	fm, err := f.browser.Form("form[id=login_form]")
	if err != nil {
		return ParseError(fUrl.String() + ": login_form")
	}

	err = fm.Input("email", email)
	if err != nil {
		return ParseError(fUrl.String() + ": login_form_email")
	}

	err = fm.Input("pass", password)
	if err != nil {
		return ParseError(fUrl.String() + ": login_form_password")
	}

	if fm.Submit() != nil {
		return ParseError(fUrl.String() + ": login_form_button")
	}

	if strings.Contains(f.browser.Title(), "Log into Facebook") {
		return AccountError(email)
	}

	return nil
}
func (f *FacebookWebservice) Logout() error {
	fUrl := buildUrl(facebooklogin)

	err := f.browser.Open(fUrl.String())
	if err != nil {
		return ConnectError(fUrl.String())
	}

	retErr := AccountError("")
	f.browser.Find("a").EachWithBreak(func(_ int, s *goquery.Selection) bool {
		if href, exist := s.Attr("href"); exist {
			if strings.Contains(href, "logout") {
				f.browser.Open(buildUrl(href).String())
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

func (f *FacebookWebservice) ChangePassword(email, oldpasswd, newpasswd string) error {
	fUrl := buildUrl(facebookchpasswd)

	err := f.browser.Open(fUrl.String())
	if err != nil {
		return ConnectError(fUrl.String())
	}

	fm, err := f.browser.Form("form[method=post]")
	if err != nil {
		return ParseError(fUrl.String() + ": chpwd_form")
	}

	err = fm.Input("password_old", oldpasswd)
	if err != nil {
		return ParseError(fUrl.String() + ": chpw_form_old")
	}

	err = fm.Input("password_new", newpasswd)
	if err != nil {
		return ParseError(fUrl.String() + ": chpw_form_new")
	}

	err = fm.Input("password_confirm", newpasswd)
	if err != nil {
		return ParseError(fUrl.String() + ": chpw_form_conf")
	}

	if fm.Submit() != nil {
		return ParseError(fUrl.String() + ": login_form_button")
	}

	// TODO put in function
	body := f.browser.Body()
	if strings.Contains(body, "Your password was incorrect") ||
		strings.Contains(body, "Password must differ from old password") {
		return ChangeError(facebookhost, email, "Password must differ from old password")
	}

	return nil
}

func (f *FacebookWebservice) GetHostname() string {
	return facebookhost
}

func buildUrl(path string) *url.URL {
	return &url.URL{
		Scheme: httpscheme,
		Host:   "m." + facebookhost,
		Path:   path,
	}
}
