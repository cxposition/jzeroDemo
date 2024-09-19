package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var C Config

type Config struct {
	Rest RestConf
	Log  LogConf

	DataBaseType string `json:",default=mysql"`
	Mysql        MysqlConf

	Banner BannerConf
}

type RestConf struct {
	rest.RestConf
}

type LogConf struct {
	logx.LogConf
}

type BannerConf struct {
	Text     string `json:",default=JZERO"`
	Color    string `json:",default=green"`
	FontName string `json:",default=starwars,options=big|larry3d|starwars|standard"`
}

type MysqlConf struct {
	Username string `json:",default=root"`
	Password string `json:",default=123456"`
	Path     string `json:",default=127.0.0.1"`
	Port     int    `json:",default=3306"`
	Dbname   string `json:",default=test"`
}
