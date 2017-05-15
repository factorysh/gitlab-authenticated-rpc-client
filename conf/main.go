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
	return os.MkdirAll(c.domainHomePath(), 0700)
}

func (c *Conf) domainHomePath() string {
	return filepath.Join(c.User.HomeDir, "."+c.Name, c.Domain)
}

func (c *Conf) tokenPath() string {
	return filepath.Join(c.domainHomePath(), "token.json")
}

func (c *Conf) SaveToken(token *oauth2.Token) error {
	err := c.makeDomainHome()
	if err != nil {
		return err
	}
	raw, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.tokenPath(), raw, 0600)
}

func (c *Conf) Token() (*oauth2.Token, error) {
	raw, err := ioutil.ReadFile(c.tokenPath())
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
