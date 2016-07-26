package models

import (
	"fmt"
	// "time"
)

func (user *User) GetCreateTime() int64 {
	return user.CreateTime
}

func (user *User) GetUpdateTime() int64 {
	return user.UpdateTime
}

func (user *User) GetNearUpdateFileTime() int64 {
	return user.NearUpdateFileTime
}

func (user *User) GetFlag() int {
	return user.Flag
}

func (user *User) GetUid() int64 {
	return user.Uid
}

func (user *User) GetWid() string {
	return user.Wid
}

func (user *User) GetUsername() string {
	return user.Username
}

func (user *User) GetPassword() string {
	return user.Password
}

func (user *User) GetTelNumber() string {
	return user.Telnumber
}

func (user *User) GetEmail() string {
	return user.Email
}

func (user *User) GetTotalConsumption() float64 {
	return user.TotalConsumption
}

func (user *User) GetFileSavePath() string {
	return user.FileSavePath
}

func (user *User) GetUploadFileNum() int64 {
	return user.UploadFileNum
}

func (user *User) GetPrintFileNum() int64 {
	return user.PrintFileNum
}

func NewUser() *User {
	user := new(User)
	return user
}

func AddUser(user *User) (err error) {
	connectDB()
	_, err = engine.Insert(user)
	if err != nil {
		return err
	}
	defer engine.Close()

	return nil
}

func CheckUser(userid int64) (has bool, err error) {
	connectDB()
	// user := &User{Uid: userid}
	// has, err = engine.Get(user)
	if err != nil {
		return false, err
	}
	defer engine.Close()

	return has, nil
}

func ModifyUser() {

}

func JudgeUser(user *User) (has bool, err error) {
	connectDB()
	has, err = engine.Get(user)
	// mt.Println(user, "judge user")
	if err != nil {
		return false, err
	}
	fmt.Println(user, "judge user")
	defer engine.Close()
	return has, nil
}
