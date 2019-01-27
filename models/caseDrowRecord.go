
package models

import (
	"paopao/db"
)


type T_case_drow_record struct {
	Id                int  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	trace_type     int
	Call_back_status  int
	Channel_id        int
	BetAmount      float32
	BeforeAmount     float32
	CurrentAmount     float32
	Create_time       string
	Remark            string
	System_order_number string
	Merchants_order_number string

}

func GetCaseDrowRecord(page int) ([]T_case_drow_record, error) {

	channel:=[]T_case_drow_record{}

	que := db.DB.Order("id desc").Offset((page - 1) * 10).Limit(10).Find(&channel)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return channel, err
}



func DeleteCaseDrowRecord(id int) (*T_case_drow_record, error) {
	channel := T_case_drow_record{Id: id}
	que := db.DB.Delete(&channel)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}

func UpdateCaseDrowRecord(id int, text string) (*T_case_drow_record, error) {
	channel := T_case_drow_record{Id: id}
	que := db.DB.Model(&channel).Update("context", text)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}