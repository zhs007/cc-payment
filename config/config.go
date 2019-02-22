package config

import (
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v2"

	"github.com/zhs007/cc-payment/err"
)

// Config - configuration
//			see the config.yaml.default
type Config struct {
	PaymentDB struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Log struct {
		LogLevel   string
		OutputType string
		LogPath    string
	}
	Service struct {
		Host string
	}
}

var cfg *Config
var onceCfg sync.Once

// LoadConfig - load config
func load(filename string) error {
	fi, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return err
	}

	cfg = &Config{}
	err = yaml.Unmarshal(fd, cfg)
	if err != nil {
		return err
	}

	err = procConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}

func isValidLogLevel(loglevel string) bool {
	return loglevel == "debug" || loglevel == "info" || loglevel == "warn" || loglevel == "error"
}

func isValidLogOutputType(outputtype string) bool {
	return outputtype == "console" || outputtype == "file"
}

// IsConsoleLog - Do you need to output to the console?
func IsConsoleLog() bool {
	if cfg == nil {
		return true
	}

	return cfg.Log.OutputType == "console"
}

func procConfig(cfg *Config) error {
	if !isValidLogLevel(cfg.Log.LogLevel) {
		return err.ErrConfigLogLevel
	}

	if !isValidLogOutputType(cfg.Log.OutputType) {
		return err.ErrConfigLogOutputType
	}

	if cfg.Service.Host == "" {
		return err.ErrInvalidConfigService
	}

	return nil
}

// LoadConfig - load config
func LoadConfig(filename string) (err error) {
	onceCfg.Do(func() {
		err = load(filename)
	})

	return
}

// GetConfig - get Config
func GetConfig() (*Config, bool) {
	if cfg == nil {
		return nil, false
	}

	return cfg, true
}
