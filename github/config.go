package github

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jingweno/gh/utils"
	"os"
	"path/filepath"
	"regexp"
)

var (
	GitHubHost string = "github.com"
)

type Config struct {
	User  string `json:"user"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

func (c *Config) FetchUser() string {
	if c.User == "" {
		var user string
		msg := fmt.Sprintf("%s username: ", c.FetchHost())
		fmt.Print(msg)
		fmt.Scanln(&user)
		c.User = user
	}

	return c.User
}

func (c *Config) FetchPassword() string {
	msg := fmt.Sprintf("%s password for %s (never stored): ", c.Host, c.User)
	fmt.Print(msg)

	pass := gopass.GetPasswd()
	if len(pass) == 0 {
		utils.Check(errors.New("Password cannot be empty"))
	}

	return string(pass)
}

func (c *Config) FetchHost() string {
	msg := fmt.Sprintf("host (%s): ", GitHubHost)
	fmt.Print(msg)
	fmt.Scanln(&c.Host)

	if c.Host == "" {
		c.Host = GitHubHost
	}

	return c.Host
}

func (c *Config) FetchTwoFactorCode() string {
	var code string
	fmt.Print("two-factor authentication code: ")
	fmt.Scanln(&code)

	return code
}

func (c *Config) FetchCredentials() {
	if c.Host == "" {
		c.FetchHost()
	}

	var changed bool
	if c.User == "" {
		c.FetchUser()
		changed = true
	}

	if c.Token == "" {
		password := c.FetchPassword()
		token, err := findOrCreateToken(c.User, password, "")
		// TODO: return an two factor auth failure error
		if err != nil {
			re := regexp.MustCompile("two-factor authentication OTP code")
			if re.MatchString(fmt.Sprintf("%s", err)) {
				code := c.FetchTwoFactorCode()
				token, err = findOrCreateToken(c.User, password, code)
			}
		}

		utils.Check(err)

		c.Token = token
		changed = true
	}

	if changed {
		err := SaveConfig(c)
		utils.Check(err)
	}
}

var (
	DefaultConfigFile = filepath.Join(os.Getenv("HOME"), ".config", "gh")
)

func CurrentConfig() *Config {
	config, err := loadConfig()
	if err != nil {
		config = Config{}
	}
	config.FetchCredentials()

	return &config
}

func loadConfig() (Config, error) {
	return loadFrom(DefaultConfigFile)
}

func loadFrom(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}

	return doLoadFrom(f)
}

func doLoadFrom(f *os.File) (Config, error) {
	defer f.Close()

	reader := bufio.NewReader(f)
	dec := json.NewDecoder(reader)

	var c Config
	err := dec.Decode(&c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func SaveConfig(config *Config) error {
	return saveTo(DefaultConfigFile, config)
}

func saveTo(filename string, config *Config) error {
	err := os.MkdirAll(filepath.Dir(filename), 0771)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	return doSaveTo(f, config)
}

func doSaveTo(f *os.File, config *Config) error {
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(config)
}

func NewConfigWithUrl(user, token, url string) Config {
	return Config{user, token, url}
}

func NewConfig(user, token string) Config {
	return Config{user, token, ""}
}
