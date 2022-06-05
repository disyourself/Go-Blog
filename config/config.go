package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Descripiton string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	//每个包里面的init都会先启动
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "Go-Blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	_, err := toml.DecodeFile("config/config.toml", &Cfg)

	if err != nil {
		panic(err)
	}
}
