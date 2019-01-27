
package models

import (
	"paopao/db"
)



type T_member_bank struct {
	Id                int  `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	User_account             string
	Bank_account_name        string
	Bank_code                string
	Bank_card_no             string
	Bank_address             string
	Creation_time            string
	Is_del                   bool
	Is_default               bool
	Trace_id                 int

}

func GetAllBankCard(page int) ([]T_member_bank, error) {

	channel:=[]T_member_bank{}

	que := db.DB.Order("id desc").Offset((page - 1) * 10).Limit(10).Find(&channel)
	if que.Error != nil {
		// panic(que.Error)
		return nil, que.Error
	}
	return channel, err
}

func AddBankCard(userAccount, bankAccountName, bankCode,bankCardNo,bankAddress,createTime string,traceId int,isDel,isDefault bool)(*T_member_bank,error)  {

	channel := T_member_bank{User_account:userAccount,Bank_account_name:bankAccountName,Bank_code:bankCode,Bank_card_no:bankCardNo,Bank_address:bankAddress,Creation_time:createTime,Is_del:isDel,Is_default:isDefault,Trace_id:traceId}

	db.DB.Create(&channel)

	return &channel, nil

}


func DeleteBankCard(id int) (*T_member_bank, error) {
	channel := T_member_bank{Id: id}
	que := db.DB.Delete(&channel)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}

func UpdateBankCard(id int, text string) (*T_member_bank, error) {
	channel := T_member_bank{Id: id}
	que := db.DB.Model(&channel).Update("context", text)
	if que.Error != nil {
		return nil, que.Error
	}
	return &channel, err
}