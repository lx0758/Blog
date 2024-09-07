package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var gConf *Conf

func Config() *Conf {
	if gConf == nil {
		content, err := os.ReadFile("application.yaml")
		if err != nil {
			fmt.Printf("Read applocation.yaml error: %s\n", err)
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
	Server     Server     `yaml:"server"`
	DataSource DataSource `yaml:"datasource"`
	Session    Session    `yaml:"session"`
}

type Server struct {
	Port    int  `yaml:"port"`
	Release bool `yaml:"release"`
}

type DataSource struct {
	Debug bool   `yaml:"debug"`
	Type  string `yaml:"type"`
	Dsn   string `yaml:"dsn"`
}

type Session struct {
	StoreKey string `yaml:"store_key"`
}
