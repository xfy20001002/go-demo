package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfigFromYaml(file string, conf interface{}) error {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		log.Println("YAML配置文件不存在：", file)
		return err
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		panic(err)
	}
	return nil
}

type Mongodb struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	DB    string `yaml:"db"`
	User  string `yaml:"user"`
	Pass  string `yaml:"pass"`
	Ok    bool   `yaml:"ok"`
	Other int64  `yaml:"other"`
}
type Mysql struct {
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	DB    string `yaml:"db"`
	User  string `yaml:"user"`
	Pass  string `yaml:"password"`
	Ok    bool   `yaml:"ok"`
	Other int64  `yaml:"other"`
}

func main() {
	type config struct {
		Mongodb Mongodb `yaml:"mongodb"`
		Mysql   Mysql   `yaml:mysql`
	}

	cfg := config{}
	cfg.Mongodb.Host = "127.0.0.1"
	cfg.Mongodb.Port = "3306"
	cfg.Mongodb.DB = "test"
	cfg.Mongodb.User = "user"
	cfg.Mongodb.Pass = "user"
	cfg.Mongodb.Ok = true
	cfg.Mongodb.Other = 123456
	LoadConfigFromYaml("yml/conf.yml", &cfg)
	fmt.Println(cfg.Mongodb)
	fmt.Println("---------------------------------------------")
	fmt.Println(cfg.Mysql)
	/*
		result:
		{localhost 3306 test test test true 123456}
		---------------------------------------------
		{localhost 3306 test test test false 0}
	*/
}
