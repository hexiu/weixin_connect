package controller

import (
	// "fmt"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/user"

	"log"
	// "time"
	"weixin_connect/models"
	"weixin_connect/modules/initConf"
)

const (
	Lang = "zh_CN"
)

var lang string = Lang

func init() {
	conf, err = initConf.InitConf()
	if err != nil {
		log.Println(err)
	}
	ok, err := conf.GetValue("UserInfo", "Lang")
	if err != nil {
		log.Println(err)
		return
	}
	lang = ok

}

func UserAddHandler(ctx *core.Context) {
	// newuser := models.NewUser()
	newuser := new(models.User)
	msg := ctx.MixedMsg
	newuser.Wid = msg.ToUserName
	newuser.OpenId = msg.FromUserName
	newuser.CreateTime = msg.CreateTime
	userinfo, err := userUpdateFromWeiXin(newuser.OpenId, "zh_CN")
	newuser.Nickname = userinfo.Nickname
	newuser.Sex = userinfo.Sex
	newuser.Country = userinfo.Country
	newuser.City = userinfo.City
	newuser.Language = userinfo.Language

	if err != nil {
		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
	}

	err = models.AddUser(newuser)
	if err != nil {
		log.Println("Controller UserHander AddUser Error : ", err)
	}
	log.Println(newuser)

}

func userUpdateFromWeiXin(openId string, lang string) (userinfo *user.UserInfo, err error) {
	userinfo, err = user.Get(Client, openId, lang)
	if err != nil {
		log.Println("controller UserHandler userUpdateFromWeiXin Error : ", err)
		return nil, err
	}
	return userinfo, nil
}
