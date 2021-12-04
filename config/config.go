package config

import (
	"log"

	"gopkg.in/ini.v1"
	"udemy.com/golang-webgosql/utils"
)

type ConfigList struct{
	Port string
	SQLDriver string
	DbName string
	LogFile string
}

var Config ConfigList

func init(){
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig(){
	cfg,err:=ini.Load("config.ini")
	if err!=nil{
		log.Fatalln(err)
	}
	Config=ConfigList{
		Port:cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
	}
}