package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

var gConf *Conf

func Config() *Conf {
	if gConf == nil {
		content, err := os.ReadFile("application.yaml")
		if err != nil {
			content = []byte(`
server:
  address: 0.0.0.0:8080
pprof:
  enable: true
  address: localhost:6060
datasource:
  type: sqlite
  path: data.db
session:
  store_key: ----------
`)
		}

		conf := &Conf{}
		err = yaml.Unmarshal(content, &conf)
		if err != nil {
			fmt.Printf("Read applocation.yaml error: %s\n", err)
		}
		gConf = conf
	}
	return gConf
}

type Conf struct {
	Pprof      Pprof      `yaml:"pprof"`
	Server     Server     `yaml:"server"`
	DataSource DataSource `yaml:"datasource"`
	Session    Session    `yaml:"session"`
}

type Pprof struct {
	Enable  bool   `yaml:"enable"`
	Address string `yaml:"address"`
}

type Server struct {
	Debug   bool   `yaml:"debug"`
	Address string `yaml:"address"`
}

type DataSource struct {
	Debug    bool   `yaml:"debug"`
	Type     string `yaml:"type"`
	Path     string `yaml:"path"`
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type Session struct {
	StoreKey string `yaml:"store_key"`
}
