package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/judger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetAllSubmit(c *gin.Context)  {
	if res := haveAuth(c, "getAllSubmit"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitModel := model.Submit{}

	submitJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			UserID 		string 	`json:"user_id" form:"user_id"`
			ProblemID 	string `json:"problem_id" form:"problem_id"`
			ContestID 	string `json:"contest_id" form:"contest_id"`
			Language 	string `json:"language" form:"language"`
			Status 		string `json:"status" form:"status"`
			MinSubmitTime 	time.Time 	`json:"min_submit_time" form:"min_submit_time"`
			MaxSubmitTime 	time.Time 	`json:"max_submit_time" form:"max_submit_time"`
		}
	}{}
	var err error
	if err = c.ShouldBind(&submitJson); err == nil {
		submitJson.Offset = (submitJson.Offset-1)*submitJson.Limit
		res := submitModel.GetAllSubmit(submitJson.Offset, submitJson.Limit, submitJson.Where.UserID,
			submitJson.Where.ProblemID, submitJson.Where.ContestID, submitJson.Where.Language,
			submitJson.Where.Status, submitJson.Where.MinSubmitTime, submitJson.Where.MaxSubmitTime)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), false))
	return
}

func GetSubmitByID(c *gin.Context) {
	if res := haveAuth(c, "getAllSubmit"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitValidate := validate.SubmitValidate
	submitModel := model.Submit{}

	var submitJson model.Submit

	if err := c.ShouldBind(&submitJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	submitMap := helper.Struct2Map(submitJson)
	if res, err:= submitValidate.ValidateMap(submitMap, "find"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := submitModel.FindSubmitByID(submitJson.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func RejudgeGroupSubmits(c *gin.Context)  {
	if res := haveAuth(c, "rejudge"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitModel := model.Submit{}

	submitJson := struct {
		UserID 		string 	`json:"user_id" form:"user_id"`
		ProblemID 	string `json:"problem_id" form:"problem_id"`
		ContestID 	string `json:"contest_id" form:"contest_id"`
		Language 	string `json:"language" form:"language"`
		Status 		string `json:"status" form:"status"`
		MinSubmitTime 	time.Time 	`json:"min_submit_time" form:"min_submit_time"`
		MaxSubmitTime 	time.Time 	`json:"max_submit_time" form:"max_submit_time"`
	}{}

	if c.ShouldBind(&submitJson) == nil {
		res := submitModel.GetSubmitGroup(submitJson.UserID, submitJson.ProblemID,
			submitJson.ContestID, submitJson.Language, submitJson.Status,
			submitJson.MinSubmitTime, submitJson.MaxSubmitTime)
		submits := res.Data.([]model.Submit)
		for _, submit := range submits {
			go func(submitData model.Submit) {
				rejudge(submitData)
			}(submit)
		}
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "发送重测成功", true))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func RejudgeSubmitByID(c *gin.Context) {
	if res := haveAuth(c, "rejudge"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	submitValidate := validate.SubmitValidate
	submitModel := model.Submit{}

	var submitJson model.Submit

	if err := c.ShouldBind(&submitJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	submitMap := helper.Struct2Map(submitJson)
	if res, err:= submitValidate.ValidateMap(submitMap, "rejudge"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := submitModel.FindSubmitByID(submitJson.ID)
	submit := res.Data.(model.Submit)
	go func(submitData model.Submit) {
		rejudge(submitData)
	}(submit)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "发送重测成功", true))
	return
}

func rejudge(submit model.Submit) {
	submitData := judger.SubmitData{
		Id: uint64(submit.ID),
		Pid: uint64(submit.ProblemID),
		Language: helper.LanguageType(submit.Language),
		Code: submit.SourceCode,
		BuildScript: "",
		RootfsConfig: nil,
	}

	callback := func(id uint64, result judger.JudgeResult) {
		// Put Result To DB
		if result.IsFinished {
			data := map[string]interface{} {
				"status": 	result.Status,
				"time": 	result.Time,
				"memory": 	result.Memory,
				"msg": 		result.Msg,
			}
			submitModel := model.Submit{}
			submitModel.UpdateStatusAfterSubmit(int(id), data)
		}
	}

	instance := judger.GetInstance()

	go instance.Submit(submitData, callback)
}