package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Cfg struct {
	Host string `default:"localhost"`
	Port string `default:"8080"`
}

const (
	appName = "avsanagrams"
)

func GetConfig() (*Cfg, error) {
	cfg := &Cfg{}
	err := envconfig.Process(appName, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
