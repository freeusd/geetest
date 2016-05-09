package geetest

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/parnurzeal/gorequest"
)

const apiHost = "api.geetest.com"

// Geetest is used for captcha registration and validation
type Geetest struct {
	captchaID  string
	privateKey string
	scheme     string
}

// New constructs and returns a Geetest
func New(captchaID, privateKey string, enableHTTPS bool) Geetest {
	scheme := "http"
	if enableHTTPS {
		scheme = "https"
	}

	return Geetest{
		captchaID:  captchaID,
		privateKey: privateKey,
		scheme:     scheme,
	}
}

// CaptchaID returns captchaID
func (g Geetest) CaptchaID() string { return g.captchaID }

// Register returns challenge get from api server
func (g Geetest) Register() (string, error) {
	_, body, errs := gorequest.New().Get(g.registrationURL()).Timeout(time.Second * 2).End()
	if errs != nil {
		return "", &multierror.Error{Errors: errs}
	}

	return hexmd5(body + g.privateKey), nil
}

// Validate validates challenge
func (g Geetest) Validate(challenge, validate, seccode string) (bool, error) {
	hash := g.privateKey + "geetest" + challenge
	if validate != hexmd5(hash) {
		return false, nil
	}

	data := fmt.Sprintf(`{"seccode":"%s"}`, seccode)
	_, body, errs := gorequest.New().Post(g.validationURL()).Query(data).End()
	if errs != nil {
		return false, &multierror.Error{Errors: errs}
	}

	return hexmd5(seccode) == body, nil
}

func (g Geetest) registrationURL() string {
	u := url.URL{}
	u.Scheme = g.scheme
	u.Host = apiHost
	u.Path = "register.php"
	query := u.Query()
	query.Set("gt", g.captchaID)
	u.RawQuery = query.Encode()
	return u.String()
}

func (g Geetest) validationURL() string {
	u := url.URL{}
	u.Scheme = g.scheme
	u.Host = apiHost
	u.Path = "validate.php"
	return u.String()
}

func hexmd5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
