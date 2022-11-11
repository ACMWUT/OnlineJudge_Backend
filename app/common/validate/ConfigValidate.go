package validate

import "OnlineJudge/app/helper"

var ConfigValidator helper.Validator

func init() {
	rules := map[string]string{
		"id":     "required",
		"name":   "required",
		"status": "required",
	}

	scenes := map[string][]string{
		"update":       {"id"},
		"changeStatus": {"id", "status"},
	}

	ConfigValidator.Rules = rules
	ConfigValidator.Scenes = scenes
}
