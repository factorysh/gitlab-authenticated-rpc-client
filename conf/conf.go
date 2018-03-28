package conf

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Conf struct {
	Name   string     // Project name, default is "gar"
	Domain string     // Server name (and maybe the port)
	User   *user.User // Current user
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
	return filepath.Join(c.domainHomePath(), "token")
}

func (c *Conf) GetToken() (string, error) {
	token := os.Getenv("GAR_TOKEN")
	if token != "" {
		return token, nil
	}
	raw_token, err := ioutil.ReadFile(c.tokenPath())
	if err == nil {
		return string(raw_token), nil
	}
	return "", nil // an empty token, it can be the first connection
}

func (c *Conf) SetToken(token string) error {
	err := c.makeDomainHome()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.tokenPath(), []byte(token), 0600)
}
