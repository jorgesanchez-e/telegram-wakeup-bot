package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	configETCLocal   = "/usr/local/etc/telegram-wakeup-bot/"
	configHome       = "$HOME/.telegram-wakeup-bot/"
	configCurrentDir = "."

	configName = "bot"
	configType = "yaml"

	confKeyWebhook = "bot.webhook"
	confKeyCert    = "bot.cert"
	confKeyCertKey = "bot.key"
	confKeyToken   = "bot.token"
)

var (
	errConfigNotFound = errors.New("config file not found")
	errNoToken        = errors.New("token not configured")
	errNoCertPath     = errors.New("certification file path not configured")
	errNoCertKeyPath  = errors.New("certification key file path not configured")
	errNoWebhook      = errors.New("telegram webhook not configured")
)

type config struct {
	webhook string
	token   string
	cert    string
	key     string
}

// New depics new config
func New() (config, error) {
	conf := config{}

	if err := conf.read(); err != nil {
		return config{}, err
	}

	return conf, nil
}

func (c *config) read() error {
	var config config

	viper.AddConfigPath(configETCLocal)
	viper.AddConfigPath(configHome)
	viper.AddConfigPath(configCurrentDir)

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errConfigNotFound
		}
		return err
	}

	if config.token = viper.GetString(confKeyToken); config.token == "" {
		return errNoToken
	}

	if config.cert = viper.GetString(confKeyCert); config.cert == "" {
		return errNoCertPath
	}

	if config.key = viper.GetString(confKeyCertKey); config.key == "" {
		return errNoCertKeyPath
	}

	if config.webhook = viper.GetString(confKeyWebhook); config.webhook == "" {
		return errNoWebhook
	}

	*c = config
	return nil
}

func (c config) Webhook() string {
	return c.webhook
}

func (c config) Cert() string {
	return c.cert
}

func (c config) Key() string {
	return c.key
}

func (c config) Token() string {
	return c.token
}
