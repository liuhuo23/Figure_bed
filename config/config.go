package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type Config struct {
	Port     int    `ini:"port"`      // 服务器端口
	SavePath string `ini:"save_path"` // 保存路径
}

func NewInit() {
	cfg, err := ini.Load("./figure.ini")
	if err != nil {
		
	}
}
