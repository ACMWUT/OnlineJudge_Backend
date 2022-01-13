package validate

import "OnlineJudge/app/helper"

var OJWebDataValidate helper.Validator

func init() {
	rules := map[string]string{
		"user_id": "required",
	}

	scenes := map[string][]string{
		"findByID":    {"id"},
		"delete":      {"user_id", "oj_name"},
		"getAll":      {"user_id", "oj_name"},
		"getLastWeek": {"user_id", "oj_name"},
		"getWeek":     {"user_id", "oj_name"},
		"getBetween":  {"user_id", "oj_name", "start_time", "end_time"},
	}

	OJWebDataValidate.Rules = rules
	OJWebDataValidate.Scenes = scenes
}
