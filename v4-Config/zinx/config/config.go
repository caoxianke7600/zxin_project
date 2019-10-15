package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//定义一个配置结构
type Config struct {
	IP         string
	Port       uint32
	Name       string
	TCPVersion string
}

//定义一个全局配置结构
var GlobalConfig Config

func LoadConfig () error {
	fmt.Println("开始读取配置文件")

	//go run server_main.go启动时，目录是基于程序运的位置
	//读取配置文件
	info,err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println("ReadFile err",err)
		return err
	}
	fmt.Println("读取的配置文件：",info)

	//解析一下配置文件，将数据放到结构体中(返回的GlobalConfig一定有&)
	err = json.Unmarshal(info,&GlobalConfig)
	if err != nil {
		fmt.Println("Unmarshal err",err)
		return err
	}
	fmt.Println("解析后：",GlobalConfig)
	return nil
}

func init()  {
	err := LoadConfig()
	if err != nil {
		fmt.Println("读取配置文件失败",err)
		os.Exit(-1)   //直接退出
	}
}
