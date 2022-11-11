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

func (model *GlobalConfig) GetAllGlobalConfig(offset int, limit int, name string) helper.ReturnType {
	var GlobalConfigs []GlobalConfig
	where := "name like ?"
	var count int

	db.Model(&GlobalConfig{}).Where(where, "%"+name+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%").
		Find(&GlobalConfigs).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"GlobalConfigs": GlobalConfigs,
				"count":         count,
			},
		}
	}
}

func (model *GlobalConfig) GetAvailGlobalConfig(name string) helper.ReturnType {
	var GlobalConfigs []GlobalConfig
	where := "status = 1 AND name like ?"

	err := db.
		Where(where, "%"+name+"%").
		Find(&GlobalConfigs).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: GlobalConfigs}
	}
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

func (model *GlobalConfig) AddGlobalConfig(newGlobalConfig GlobalConfig) helper.ReturnType { //jun
	GlobalConfig := GlobalConfig{}

	if err := db.Where("name = ?", newGlobalConfig.Name).First(&GlobalConfig).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "标签已存在", Data: false}
	}

	err := db.Create(&newGlobalConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *GlobalConfig) DeleteGlobalConfig(GlobalConfigID int) helper.ReturnType {
	err := db.Where("id = ?", GlobalConfigID).Delete(GlobalConfig{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *GlobalConfig) UpdateGlobalConfig(GlobalConfigID int, updateGlobalConfig GlobalConfig) helper.ReturnType {
	err := db.Model(&GlobalConfig{}).Where("id = ?", GlobalConfigID).Update(updateGlobalConfig).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *GlobalConfig) ChangeGlobalConfigStatus(GlobalConfigID int, status int) helper.ReturnType {
	err := db.Model(&GlobalConfig{}).Where("id = ?", GlobalConfigID).Update("status", status).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}
