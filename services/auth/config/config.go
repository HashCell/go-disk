package config

import (
	"fmt"
	"go-disk/common/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

var Conf *Config

func init() {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println(err)
	}
	ct := reflect.TypeOf(*Conf)
	elements := reflect.ValueOf(Conf)
	utils.SetConfigDefaultValue(ct, elements)
	log.Println(Conf)
}

type Config struct {

	//db
	DataSource struct {
		Mysql struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Host string `yaml:"host"`
			Port int `yaml:"port"`
			TimeLoc string `yaml:"timeLoc"`
			Database string `yaml:"database"`
			MaxIdle int `yaml:"maxIdle" default:"10"`
			MaxOpenConn int `yaml:"maxOpenConn" default:"100"`
			MaxLifeTime int `yaml:"maxLifeTime" default:"720"`
		} `yaml:"mysql"`

		Redis struct {
			Addr string `yaml:"addr"`
			Password string `yaml:"password"`
			Database int `yaml:"database"`
			TokenExpireTime int `yaml:"tokenExpireTime"`
		} `yaml:"redis"`
	} `yaml:"dataSource"`

	Business struct {
		UserPasswordSalt string `yaml:"userPasswordSalt"`
	} `yaml:"business"`

	Micro struct {
		Registration struct {
			Consul struct {
				Addr string `yaml:"addr"`
			} `yaml:"consul"`
		} `yaml:"registration"`
	} `yaml:"micro"`
}



