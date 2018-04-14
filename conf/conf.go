package conf

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type Conf struct {
	Name       string     // Project name, default is "gar"
	Domain     string     // Server name (and maybe the port)
	User       *user.User // Current user
	token      string
	WriteToken bool
}

func NewConf(name, domain string) *Conf {
	usr, _ := user.Current()
	return &Conf{
		Name:       name,
		Domain:     domain,
		User:       usr,
		WriteToken: true,
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

// GetToken get the JWT token from file or ENV, "" when no token is available
func (c *Conf) GetToken() (string, error) {
	if c.token != "" {
		return c.token, nil
	}
	token := os.Getenv("GAR_TOKEN")
	if token != "" {
		log.WithFields(log.Fields{
			"ENV":   "GAR_TOKEN",
			"token": token,
		}).Info("GetToken")
		c.token = token
		return token, nil
	}
	path := c.tokenPath()
	rawToken, err := ioutil.ReadFile(path)
	if err == nil {
		log.WithFields(log.Fields{
			"path":  path,
			"token": string(rawToken),
		}).Info("GetToken")
		c.token = string(rawToken)
		return string(rawToken), nil
	}
	return "", nil // an empty token, it can be the first connection
}

func (c *Conf) SetToken(token string) error {
	c.token = token
	if c.WriteToken {
		err := c.makeDomainHome()
		if err != nil {
			return err
		}
		return ioutil.WriteFile(c.tokenPath(), []byte(token), 0600)
	}
	return nil
}
