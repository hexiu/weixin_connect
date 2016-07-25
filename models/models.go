package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"zypc_submit/modules"
)

const (
	DataType = "mysql"
	Database = "infomation"
	Username = "root"
	Password = "axiu"
	Host     = "127.0.0.1"
	Port     = "3306"
)

var (
	datatype string
	database string
	username string
	password string
	host     string
	port     string
)

var engine *xorm.Engine

type User struct {
	UserId    int64 `xorm:"index"`
	UserName  string
	Password  string
	Time      time.Time `xorm:"index"`
	Email     string
	Telnumber string
	Flag      int
}

type Topic struct {
	UserId  int64 `xorm:"index"`
	Content string
	Flag    int
	Time    time.Time `xorm:"index"`
}

type Infomation struct {
	UserId int64 `xorm:"index"`
	Flag   int64
	Info1  string
	Info2  string
	Info3  string
	Info4  string
	Info5  string
	Info6  string
	Info7  string
	Info8  string
	Info9  string
	Info10 string
	Info11 string
	Info12 string
	Info13 string
	Info14 string
	Info15 string
	Info16 string
	Info17 string
	Info18 string
	Info19 string
	Info20 string
	Time   time.Time `xorm:"index"`
}

func init() {
	err := initconf()
	if err != nil {
		fmt.Println(err)
	}

}

func initconf() (err error) {
	conf, err := modules.InitConf()
	if err != nil {
		return err
	}

	fmt.Println(conf)

	if ok, err := conf.GetValue("DataControl", "DataType"); err == nil {
		datatype = ok
	} else {
		datatype = DataType
	}

	if ok, err := conf.GetValue("DataControl", "DataBase"); err == nil {
		database = ok
	} else {
		database = Database
	}
	if ok, err := conf.GetValue("DataControl", "Username"); err == nil {
		username = ok
	} else {
		username = Username
	}
	if ok, err := conf.GetValue("DataControl", "Password"); err == nil {
		password = ok
	} else {
		password = Password
	}
	if ok, err := conf.GetValue("DataControl", "Host"); err == nil {
		host = ok
	} else {
		host = Host
	}
	if ok, err := conf.GetValue("DataControl", "Port"); err == nil {
		port = ok
	} else {
		port = Port
	}

	return nil
}

func connectDB() (err error) {
	engine, err = xorm.NewEngine(datatype, username+":"+password+"@tcp("+host+":"+port+")"+"/"+database+"?charset=utf8")
	if err != nil {
		return err
	}
	return nil

}

func RegisterDB() (err error) {
	err = connectDB()
	if err != nil {
		return err
	}

	fmt.Println(engine.Ping())

	if ok, _ := engine.IsTableExist("user"); !ok {
		engine.CreateTables(new(User))
	}

	if ok, _ := engine.IsTableExist("topic"); !ok {
		engine.CreateTables(new(Topic))
	}

	if ok, _ := engine.IsTableExist("infomation"); !ok {
		engine.CreateTables(new(Infomation))
	}

	defer engine.Close()
	return nil
}
