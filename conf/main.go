package conf

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Conf struct {
	Name   string
	Domain string
	User   *user.User
}

func NewConf(name, domain string) *Conf {
	usr, _ := user.Current()
	return &Conf{
		Name:   name,
		Domain: domain,
		User:   usr,
	}
}

func (c *Conf) makeDomainHome() error {
	domainPath := filepath.Join(c.User.HomeDir, "."+c.Name, c.Domain)
	return os.MkdirAll(domainPath, 0700)
}

func (c *Conf) Token() (*oauth2.Token, error) {
	tokenPath := filepath.Join(c.User.HomeDir, "."+c.Name, c.Domain,
		"token.json")
	raw, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		if os.IsNotExist(err) { // file not exist
			return nil, nil
		}
		return nil, err
	}
	var token oauth2.Token
	err = json.Unmarshal(raw, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
