
package models

import (
	"paopao/db"
)


type T_case_flow_record struct {
	Id                int  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Record_status     int
	Call_back_status  int
	Channel_id        int
	Channel_name      string
	Order_amount      float32
	Commit_amount     float32
	Rebate_amount     float32
	Create_time       string
	Remark            string
	System_order_number string
	Merchants_order_number string

}

func GetCaseFlowRecord(page int) ([]T_case_flow_record, error) {

	channel:=[]T_case_flow_record{}

	que := db.DB.Order("id desc").Offset((page - 1) * 10).Limit(10).Find(&channel)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return channel, err
}



func DeleteCaseFlowRecord(id int) (*T_case_flow_record, error) {
	channel := T_case_flow_record{Id: id}
	que := db.DB.Delete(&channel)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}

func UpdateCaseFlowRecord(id int, text string) (*T_case_flow_record, error) {
	channel := T_case_flow_record{Id: id}
	que := db.DB.Model(&channel).Update("context", text)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}