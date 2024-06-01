package core

import (
	"admin/configs"
	"admin/global"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// InitConfig 读取yaml文件配置
func InitConfig() {
	const ConfigFile = "./configs/settings.yaml"
	c := &configs.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取yaml文件错误: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("yaml文件序列化错误: %v", err)
	}
	log.Println("yaml文件初始化成功")
	global.Config = c
}
