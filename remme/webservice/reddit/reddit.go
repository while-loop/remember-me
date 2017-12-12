package reddit

import (
	"net/url"
	"strings"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	"github.com/while-loop/remember-me/remme/log"
	"github.com/while-loop/remember-me/remme/util"
	"github.com/while-loop/remember-me/remme/webservice"
)

const (
	httpScheme     = "https"
	redditHost     = "reddit.com"
	redditChPasswd = "/prefs/update/"
	redditLogin    = redditChPasswd
	redditLogout   = "/logout"
)

func init() {
	webservice.Register(redditHost, New())
}

type redditWebservice struct {
}

func New() webservice.Webservice {
	return &redditWebservice{}
}

func (f *redditWebservice) login(browsr *browser.Browser, email, password string) error {
	fUrl := buildUrl(redditLogin)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservice.ConnectError{Hostname: fUrl.String()}
	}

	// Log in to the site.
	fm, err := browsr.Form("form[id=login-form]")
	if err != nil {
		log.Error(err)
		return webservice.ParseError{Hostname: fUrl.String() + ": login-form"}
	}

	err = fm.Input("user", email)
	if err != nil {
		log.Error(err)
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_username"}
	}

	err = fm.Input("passwd", password)
	if err != nil {
		log.Error(err)
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_password"}
	}

	if fm.Submit() != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_button"}
	}

	body := browsr.Body()
	if strings.Contains(body, "incorrect username or password") {
		return webservice.AccountError{Email: email, Hostname: redditHost}
	}

	return nil
}

func (f *redditWebservice) logout(browsr *browser.Browser) (retErr error) {
	retErr = nil
	err := browsr.Post(buildUrl(redditLogout).String(), "plain/html", nil)
	if err != nil {
		log.Error("Failed to log out of reddit", err)
		retErr = &webservice.AccountError{Hostname: redditHost}
	}

	return retErr
}

func (f *redditWebservice) changePassword(browsr *browser.Browser, email, oldpasswd, newpasswd string) error {
	fUrl := buildUrl(redditChPasswd)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservice.ConnectError{Hostname: fUrl.String()}
	}

	body := browsr.Body()
	log.Debug(body)
	fm, err := browsr.Form("form[id=pref-update-password]")
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpwd_form"}
	}

	err = fm.Input("curpass", oldpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_curpass"}
	}

	err = fm.Input("newpass", newpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_newpass"}
	}

	err = fm.Input("verpass", newpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_verpass"}
	}

	fm.Set("submit", "save")
	if err = fm.Submit(); err != nil {
		log.Error("Unable to find chpwd submit button", err)
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_button"}
	}

	// TODO put in function
	body = browsr.Body()
	err = nil
	if strings.Contains(body, "error WRONG_PASSWORD field-curpass") {
		err = webservice.AccountError{
			Email:    email,
			Hostname: redditHost,
		}
	} else if strings.Contains(body, "error OLD_PASSWORD_MATCH field-newpass") {
		err = webservice.ChangeError{
			Hostname: redditHost,
			Email:    email,
			Message:  "Password must differ from old password",
		}
	} else if strings.Contains(body, "error BAD_PASSWORD_MATCH field-verpass") {
		err = webservice.ChangeError{
			Hostname: redditHost,
			Email:    email,
			Message:  "New passwords do not match",
		}
	}
	return err
}

func (f *redditWebservice) ChangePassword(email, oldpasswd, newpasswd string) error {
	browsr := surf.NewBrowser()
	se := util.StickyError{}

	se.Swallow(f.login(browsr, email, oldpasswd))
	se.Swallow(f.changePassword(browsr, email, oldpasswd, newpasswd))
	se.Swallow(f.logout(browsr))

	return se.Error()
}

func buildUrl(path string) *url.URL {
	return &url.URL{
		Scheme: httpScheme,
		Host:   redditHost,
		Path:   path,
	}
}
