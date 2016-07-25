package main

import (
	// "fmt"
	"github.com/Unknwon/goconfig"
	// "github.com/chanxuehong/wechat.v2/mp/core"
	// "github.com/chanxuehong/wechat.v2/mp/menu"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	// "github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"log"
	"net/http"
	"weixin_connect/controller"
	"weixin_connect/models"
	"weixin_connect/modules/initConf"
)

// 默认变量值
const (
	Port = "8080"
)

//  初始化配置相关的变量
var (
	conf *goconfig.ConfigFile
	err  error
	port string = Port
)

func init() {
	conf, err = initConf.InitConf()
	if err != nil {
		log.Println(err)
	}
	initconf()
	models.RegisterDB()

	http.HandleFunc("/wx_callback", controller.WxCallbackHandler)
	// controller.MenuHandler()
	http.HandleFunc("/page1", controller.OauthHandler)
	http.HandleFunc("/page2", controller.Page2Handler)

}

func initconf() {
	if ok, err := conf.GetValue("Server", "ListenPort"); err == nil {
		port = ok
	} else {
		log.Printf("[Wranning :]  Use Default Port %s ", Port)
	}

}

func main() {
	log.Println(http.ListenAndServe(":"+port, nil))
}
