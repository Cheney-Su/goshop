package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
)

func SetUpDefaultConfig() {
	config := flag.String("default", "", "SetUpDefaultConfig")
	flag.Parse()
	if *config == "" {
		*config = "./src/resources/default"
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName(*config)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("读取默认配置信息失败:", err.Error())
	} else {
		log.Println("成功读取默认配置信息...")
	}
}

func SetUpStatic() {
	iris.Static("/css", "./src/static/css", 1)
	iris.Static("/image", "./src/static/image", 1)
	iris.Static("/js", "./src/static/js", 1)
	iris.Static("/views", "./src/views", 1)
}

func SetUpHtmlTemplates() {
	iris.UseTemplate(html.New()).Directory("./src/views", ".html")
}
