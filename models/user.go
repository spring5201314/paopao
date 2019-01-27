
package models

import (
	"paopao/db"
	"time"
)

type T_merchants_user struct {
	Id      int    `gorm:"AUTO_INCREMENT" `
	Trace_id  int
	Trace_key string
	Phone_number int
	User_name string
	Pass_word string
	Sex      int
	Corporation_name string
	Id_number string
	Business_address string
	Email string
	Qq string
	RegTime  string
}

func UserRegister(name, pass,tracekey,corporationName,idNumber,businessAddress,email,qq,regTime string,traceId,phoneNum,sex int) (*T_merchants_user, error) {
	t := time.Now()
	regTime = t.Format("2006-01-02 15:04:05")
	// user := User{UserName: name, PassWord: pass, RegTime: t.Format("2006-01-02 15:04:05")}
	user := T_merchants_user{User_name: name, Pass_word: pass,Trace_key:tracekey,Corporation_name:corporationName,Id_number:idNumber,Business_address:businessAddress,Email:email,Qq:qq,RegTime:regTime,Trace_id:traceId,Phone_number:phoneNum, Sex: sex}
	db.DB.Create(&user)
	return &user, nil
}

func UserLogin(name, pass string) (*T_merchants_user, error) {
	user := T_merchants_user{}
	var err error
	que := db.DB.Where("user_name = ? AND pass_word = ?", name, pass).First(&user)
	if que.Error != nil {
		return nil, err
	}
	if len(user.User_name) != 0 {
		return &user, nil
	}

	return nil, err
}

func GetAllUser(page int) ([]T_merchants_user, error) {

	channel:=[]T_merchants_user{}

	que := db.DB.Order("id desc").Offset((page - 1) * 10).Limit(10).Find(&channel)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return channel, err
}

func GetUser(uid int) (*T_merchants_user, error) {
	user := T_merchants_user{}
	que := db.DB.Where("id = ?", uid).Find(&user)
	if que.Error != nil {
		// panic(que.Error)
		return nil, err
	}
	return &user, nil
}

func GetName(name string) bool {
	user := T_merchants_user{}
	if err := db.DB.Where("user_name = ?", name).Find(&user).Error; err != nil {
		// panic(que.Error)
		return false
	}
	if len(user.User_name) != 0 {
		return true
	}
	return false
}


func DeleteUser(id int) (*T_merchants_user, error) {
	channel := T_merchants_user{Id: id}
	que := db.DB.Delete(&channel)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}

func UpdateUser(id int, text string) (*T_merchants_user, error) {
	channel := T_merchants_user{Id: id}
	que := db.DB.Model(&channel).Update("context", text)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}
