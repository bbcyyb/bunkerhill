package config

import (
	"log"
	"sync"
)

var (
	instance = &Adapter{make(map[string]*Config), make([]string)}
	mutex    = sync.Mutex
)

type Adapter struct {
	confs    map[string]*Config
	sequence []string
}

type Config interface {
	GetValue(string) string
	Load() error
}

func NewAdapter() *Adapter {
	return instance
}

func (a *Adapter) Register(mode, path string) {
	log.Printf("** Register configuration module [%s]\n", mode)
	mutex.Lock()
	if _, ok := a.confs[mode]; ok {
		panic("config: Register called twice for adapter " + name)
	}

	conf, err := newConfig(mode, path)
	if err != nil {
		panic("config: Attempt to create new config entry " + mode + " failed.")
	}

	a.confs[mode] = conf
	a.sequence = Append(a.sequence, mode)

	mutex.Unlock()
}

func (a *Adapter) GetValue(key string) string {
	if len(a.confs) == 0 {
		return ""
	}

	for _, mode := range a.confs {
		if result := a.confs[mode].GetValue(key); result != "" {
			return result
		}
	}
	return ""
}

func newConfig(mode, path string) (*Config, error) {
	var newConf *Config
	switch mode {
	case "env":
		newConf = environ.NewEnvConfig()
	case "ini":
		newConf = ini.NewIniConfig(path)
	default:
		panic("Don't supported configure type.")
	}

	if err := newConf.Load(); err != nil {
		return nil, err
	}
	return newConf, err
}
