package models

import (
	// "fmt"
	"log"
	// "time"
)

func AddImageInfo(fileinfo *FileInfo) (err error) {
	connectDB()
	_, err = engine.Insert(fileinfo)
	if err != nil {
		return err
	}
	defer engine.Close()

	return nil
}

func GetFileInfo(openid, wid string) (fileinfo *FileInfo, err error) {
	connectDB()
	fileinfo = &FileInfo{
		Wid:    wid,
		OpenId: openid,
	}
	has, err := engine.Get(fileinfo)
	if err != nil {
		return nil, err
	}
	if has == true {
		return fileinfo, nil
	} else {
		log.Println("no file!!!", wid, openid)
		return nil, err
	}
	defer engine.Close()
	return nil, err
}
