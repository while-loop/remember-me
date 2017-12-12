package facebook

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"github.com/while-loop/remember-me/remme/util"
	"github.com/while-loop/remember-me/remme/webservice"
)

const (
	httpscheme       = "https"
	facebookhost     = "facebook.com"
	facebooklogin    = "/login.php"
	facebookchpasswd = "/settings/security/password/"
)

func init() {
	webservice.Register(facebookhost, New())
}

type facebookWebservice struct {
}

func New() webservice.Webservice {
	return &facebookWebservice{}
}

func (f *facebookWebservice) login(browsr *browser.Browser, email, password string) error {
	fUrl := buildFBUrl(facebooklogin)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservice.ConnectError{Hostname: fUrl.String()}
	}

	// Log in to the site.
	fm, err := browsr.Form("form[id=login_form]")
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form"}
	}

	err = fm.Input("email", email)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_email"}
	}

	err = fm.Input("pass", password)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_password"}
	}

	if fm.Submit() != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_button"}
	}

	if strings.Contains(browsr.Title(), "Log into Facebook") {
		return webservice.AccountError{Email: email, Hostname: facebookhost}
	}

	return nil
}

func (f *facebookWebservice) logout(browsr *browser.Browser) error {
	fUrl := buildFBUrl(facebooklogin)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservice.ConnectError{Hostname: fUrl.String()}
	}

	retErr := &webservice.AccountError{Hostname: facebookhost}
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

func (f *facebookWebservice) changePassword(browsr *browser.Browser, email, oldpasswd, newpasswd string) error {
	fUrl := buildFBUrl(facebookchpasswd)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservice.ConnectError{Hostname: fUrl.String()}
	}

	fm, err := browsr.Form("form[method=post]")
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpwd_form"}
	}

	err = fm.Input("password_old", oldpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_old"}
	}

	err = fm.Input("password_new", newpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_new"}
	}

	err = fm.Input("password_confirm", newpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_conf"}
	}

	if fm.Submit() != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_button"}
	}

	// TODO put in function
	body := browsr.Body()
	err = nil
	if strings.Contains(body, "password was incorrect") {
		err = webservice.AccountError{
			Email:    email,
			Hostname: facebookhost,
		}
	} else if strings.Contains(body, "Password must differ from old password") {
		err = webservice.ChangeError{
			Hostname: facebookhost,
			Email:    email,
			Message:  "Password must differ from old password",
		}
	}
	return err
}

func (f *facebookWebservice) ChangePassword(email, oldpasswd, newpasswd string) error {
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
