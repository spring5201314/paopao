
package models

import (
	"paopao/db"
)

var err error


type T_channel struct {
	Id                int  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Channel_code             string
	Channel_name             string
	Channel_status           string
	Rate              float32

}

func GetAllChannel(page int) ([]T_channel, error) {

	channel:=[]T_channel{}

	que := db.DB.Order("id desc").Offset((page - 1) * 10).Limit(10).Find(&channel)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return channel, err
}

func AddChannel(Channel_code, Channel_name, Channel_status string,Rate float32)(*T_channel,error)  {

	channel := T_channel{Channel_code:Channel_code,Channel_name:Channel_name,Channel_status:Channel_status,Rate:Rate}

	db.DB.Create(&channel)

	return &channel, nil

}


func DeleteChannel(id int) (*T_channel, error) {
	channel := T_channel{Id: id}
	que := db.DB.Delete(&channel)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}

func UpdateChannel(id int, text string) (*T_channel, error) {
	channel := T_channel{Id: id}
	que := db.DB.Model(&channel).Update("context", text)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}