package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	// "time"
	"weixin_connect/modules/initConf"
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
	Uid                int64 `xorm:"index"`
	Wid                string
	OpenId             string
	Username           string
	Password           string
	Email              string
	Telnumber          string
	TotalConsumption   float64
	FileSavePath       string
	UploadFileNum      int64
	PrintFileNum       int64
	CreateTime         int64 `xorm:"index"`
	UpdateTime         int64 `xorm:"index"`
	NearUpdateFileTime int64 `xorm:"index"`
	Flag               int
	Nickname           string
	Sex                int
	Language           string
	City               string
	Province           string
	Country            string
}

type FileInfo struct {
	Fid            int64 `xorm:"index"`
	Wid            string
	OpenId         string
	FileWherePath  string //标识文件存储位置：在互联网还是微信端。
	FileName       string
	FileReName     string
	FileUploadDate string
	FilePrintTime  int64
	FilePayInfo    bool
	FileType       string
	MediaId        string
	MsgId          int64
	FileUrl        string
	Flag           int
	FileUploadTime int64 `xorm:"index"`
}

type PayInfo struct {
	Zid           int64 `xorm:"index"`
	Wid           string
	OpenId        string
	PrintFile     string
	PrintFiletype string
	PrintFileUrl  string
	PrintFiletime string
	PayMoney      float64
	PayTime       int64 `xorm:"index"`
	CreateTime    int64 `xorm:"index"`
	PrintOk       bool
	PayOk         bool
}

func init() {
	err := initconf()
	if err != nil {
		fmt.Println(err)
	}

}

func initconf() (err error) {
	conf, err := initConf.InitConf()
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

	if ok, _ := engine.IsTableExist("User"); !ok {
		engine.CreateTables(new(User))
	}

	if ok, _ := engine.IsTableExist("FileInfo"); !ok {
		engine.CreateTables(new(FileInfo))
	}

	if ok, _ := engine.IsTableExist("PayInfo"); !ok {
		engine.CreateTables(new(PayInfo))
	}

	defer engine.Close()
	return nil

}
