package conf

import (
	"github.com/satori/go.uuid"
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

func (c *Conf) idPath() string {
	return filepath.Join(c.domainHomePath(), "user.id")
}

func (c *Conf) Token() (string, error) {
	id, err := ioutil.ReadFile(c.idPath())
	if err == nil {
		return string(id), nil
	}
	if os.IsNotExist(err) { // file not exist
		err = c.makeDomainHome()
		if err != nil {
			return "", err
		}
		u := uuid.NewV4()
		err = ioutil.WriteFile(c.idPath(), u.Bytes(), 0600)
		if err != nil {
			return "", err
		}
		return u.String(), nil
	}
	return "", err
}
