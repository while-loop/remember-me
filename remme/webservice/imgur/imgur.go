package imgur

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
	httpScheme    = "https"
	imgurHost     = "imgur.com"
	imgurChPasswd = "/account/settings/password"
	imgurLogin    = "/signin"
	imgurLogout   = "/logout"
)

func init() {
	webservice.Register(imgurHost, New())
}

type imgurWebservice struct {
}

func New() webservice.Webservice {
	return &imgurWebservice{}
}

func (f *imgurWebservice) login(browsr *browser.Browser, email, password string) error {
	fUrl := buildUrl(imgurLogin)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservice.ConnectError{Hostname: fUrl.String()}
	}

	// Log in to the site.
	fm, err := browsr.Form("form[id=signin-form]")
	if err != nil {
		log.Error(err)
		return webservice.ParseError{Hostname: fUrl.String() + ": login-form"}
	}

	err = fm.Input("username", email)
	if err != nil {
		log.Error(err)
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_username"}
	}

	err = fm.Input("password", password)
	if err != nil {
		log.Error(err)
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_password"}
	}

	if fm.Submit() != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": login_form_button"}
	}

	body := browsr.Body()
	if strings.Contains(body, "fill out a captcha") {
		return webservice.ChangeError{Email: email, Hostname: imgurHost, Message: "captcha required"}
	} else if !strings.HasPrefix(strings.TrimSpace(browsr.Title()), "Imgur:") {
		return webservice.AccountError{Email: email, Hostname: imgurHost}
	}

	return nil
}

func (f *imgurWebservice) changePassword(browsr *browser.Browser, email, oldpasswd, newpasswd string) error {
	fUrl := buildUrl(imgurChPasswd)

	err := browsr.Open(fUrl.String())
	if err != nil {
		return webservice.ConnectError{Hostname: fUrl.String()}
	}

	body := browsr.Body()
	log.Debug(body)
	fm, err := browsr.Form("form[id=password-form]")
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpwd_form"}
	}

	err = fm.Input("current-password", oldpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_curpass"}
	}

	err = fm.Input("password", newpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_newpass"}
	}

	err = fm.Input("confirm_password", newpasswd)
	if err != nil {
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_verpass"}
	}

	if err = fm.Submit(); err != nil {
		log.Error("Unable to find chpwd submit button", err)
		return webservice.ParseError{Hostname: fUrl.String() + ": chpw_form_button"}
	}

	// TODO put in function
	body = browsr.Body()
	err = nil
	if !strings.Contains(body, "Your account has been updated") {
		err = webservice.AccountError{
			Email:    email,
			Hostname: imgurHost,
		}
	}
	return err
}

func (f *imgurWebservice) logout(browsr *browser.Browser) (retErr error) {
	retErr = nil
	err := browsr.Post(buildUrl(imgurLogout).String(), "plain/html", nil)
	if err != nil {
		log.Error("Failed to log out of imgur", err)
		retErr = &webservice.AccountError{Hostname: imgurHost}
	}

	return retErr
}

func (f *imgurWebservice) ChangePassword(email, oldpasswd, newpasswd string) error {
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
		Host:   imgurHost,
		Path:   path,
	}
}
