package controller

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"log"
	"net/http"
	// "strconv"
	"weixin_connect/modules/initConf"
)

var (
	// 微信对接相关变量
	WxAppId         string = ""
	WxAppSecret     string = ""
	WxOriId         string = ""
	WxToken         string
	WxEncodedAESKey string = ""
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server
)

var (
	conf *goconfig.ConfigFile
	err  error
)

func init() {

	conf, err = initConf.InitConf()
	if err != nil {
		log.Println(err)
	}
	initconf()

	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	msgHandler = mux
	msgServer = core.NewServer(WxOriId, WxAppId, WxToken, WxEncodedAESKey, msgHandler, nil)
	fmt.Println(msgServer)

}

func initconf() {
	if ok, err := conf.GetValue("WeiXin", "WxAppId"); err == nil {
		WxAppId = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("WeiXin", "WxAppSecret"); err == nil {
		WxAppSecret = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("WeiXin", "WxOriId"); err == nil {
		WxOriId = ok
	} else {
		log.Println(err)
	}
	if ok, err := conf.GetValue("WeiXin", "WxToken"); err == nil {
		if ok == "" {
			log.Println("WeiXin Config Error : ", "WxToken can not null")
			return
		} else {
			WxToken = ok
			fmt.Println(WxToken)
		}

	} else {
		log.Println(err)
		return
	}
	if ok, err := conf.GetValue("WeiXin", "WxEncodedAESKey"); err == nil {
		WxEncodedAESKey = ok
	} else {
		log.Println(err)
	}

}

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func WxCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// MenuHandler()
	// CreateMenu()
	// menuCreateHandler()
	msgServer.ServeHTTP(w, r, nil)

}

func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)
	fmt.Println(msg)
	msg.Content = fmt.Sprintln(msg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp) // 明文回复
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	ProvideQrcode()
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	msg := ctx.MixedMsg
	msg.Content = fmt.Sprintln(msg)
	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp)
	AddMediaInfo(ctx)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)
	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	ctx.RawResponse(resp) // 明文回复
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	// DeleteMenu()

}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	if ctx.MixedMsg.EventType == "subscribe" {
		UserAddHandler(ctx)
	}

	ctx.NoneResponse()
}
