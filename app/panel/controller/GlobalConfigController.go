package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// =========
// save to redis
// status
// 0/false for disable
// 1/true for enable
// =========
// config
// Register
// UpdateUserInfo
// =========
// Key:
// {env}:GC:Register
// {env}:GC:UpdateUserInfo
// =========

func UpdateConfigStatus(c *gin.Context) {
	configValidator := validate.ConfigValidator
	gconfigModel := model.GlobalConfig{}

	var gconfigJson model.GlobalConfig
	if err := c.ShouldBind(&gconfigJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	log.Println(gconfigJson)
	gcMap := helper.Struct2Map(gconfigJson)
	if res, err := configValidator.ValidateMap(gcMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := gconfigModel.ChangeGlobalConfigStatus(gconfigJson.ID, gconfigJson.Status)

	// Update Status In Redis
	var redisStr string
	// TODO: refactor to Map -> redis string
	if gconfigJson.ID == 1 {
		redisStr = redis_key.RegisterPermission()
	} else {
		redisStr = redis_key.UpdateUserInfoPermission()
	}
	_ = database.PutToRedis(redisStr, gconfigJson.Status, 3600)

	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetAllConfigStatus(c *gin.Context) {
	tagModel := model.GlobalConfig{}

	tagJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			Name string `json:"name" form:"name"`
		}
	}{}

	if c.ShouldBind(&tagJson) == nil {
		tagJson.Offset = (tagJson.Offset - 1) * tagJson.Limit
		res := tagModel.GetAllGlobalConfig(tagJson.Offset, tagJson.Limit, tagJson.Where.Name)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return

}
