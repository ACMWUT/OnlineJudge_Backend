package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type GlobalConfig struct {
	ID     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Status int    `json:"status" form:"status"`
}

func (model *GlobalConfig) FindGlobalConfigByID(id int) helper.ReturnType {
	var GlobalConfig GlobalConfig

	err := db.Where("id = ?", id).First(&GlobalConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: GlobalConfig}
	}
}
