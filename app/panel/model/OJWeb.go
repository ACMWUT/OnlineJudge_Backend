package model

//OJWeb 用于储存oj网站的数据表
type OJWeb struct {
	OJID   int    `json:"oj_id" form:"oj_id"`
	OJName string `json:"oj_name" form:"oj_name"`
	OJUrl  string `json:"oj_url" form:"oj_url"`
	Status int    `json:"status" form:"status"`
}

//TableName 设定表名
func (OJWeb) TableName() string {
	return "oj_webs"
}

//TODO:delete this later
// func (model *User) SetAdmin(userID int, isAdmin int) helper.ReturnType {
// 	err := db.Model(&User{}).Where("user_id = ?", userID).Update("identity", isAdmin).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
// 	}
// }

// func (model *User) UpdateUser(userID int, updateUser User) helper.ReturnType {
// 	err := db.Model(&User{}).Where("user_id = ?", userID).Update(updateUser).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
// 	}
// }

// func (model *User) DeleteUser(userID int) helper.ReturnType {
// 	err := db.Where("user_id = ?", userID).Delete(User{}).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
// 	}
// }

// //AddUser 添加用户
// func (model *User) AddUser(newUser User) helper.ReturnType {
// 	user := User{}

// 	if err := db.Where("nick = ? OR mail = ?", newUser.Nick, newUser.Mail).First(&user).Error; err == nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "昵称或邮箱已存在", Data: false}
// 	}

// 	err := db.Create(&newUser).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
// 	}
// }

// //AddUsersAndContestUsers 添加多个用户和比赛用户，不检测邮箱冲突
// //此处对contestUser也进行了添加
// func (model *User) AddUsersAndContestUsers(newUsers []User, contestID int) helper.ReturnType {
// 	user := User{}
// 	contestUserJSON := struct {
// 		ContestID int `json:"contest_id" form:"contest_id"`
// 	}{}
// 	tx := db.Begin()
// 	for _, newUser := range newUsers {
// 		if err := tx.Where("nick = ?", newUser.Nick).First(&user).Error; err == nil {
// 			tx.Rollback()
// 			return helper.ReturnType{Status: constants.CodeError, Msg: "昵称已存在", Data: false}
// 		}
// 		err := tx.Create(&newUser).Error
// 		if err != nil {
// 			tx.Rollback()
// 			return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
// 		}
// 		var contestUser ContestUser

// 		contestUser.ContestID = contestUserJSON.ContestID

// 		findUser := user.GetUserByNick(user.Nick)
// 		if findUser.Status != constants.CodeSuccess {
// 			tx.Rollback()
// 			return helper.ReturnType{Status: constants.CodeError, Msg: "数据库错误", Data: false}
// 		}
// 		user_to_id := (findUser.Data).(User)
// 		contestUser.UserID = user_to_id.UserID
// 		if err := tx.Where("contest_id = ?", contestUser.ContestID).Where("user_id = ?", contestUser.UserID).Find(&contestUser).Error; err == nil {
// 			tx.Rollback()
// 			return helper.ReturnType{Status: constants.CodeError, Msg: "已经参加比赛，请勿重复参赛", Data: ""}
// 		}

// 		err = tx.Create(&contestUser).Error
// 		if err != nil {
// 			tx.Rollback()
// 			return helper.ReturnType{Status: constants.CodeError, Msg: "参加比赛失败", Data: ""}
// 		}
// 	}
// 	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}

// }

// func (model *User) CheckLogin(loginUser User) helper.ReturnType {
// 	user := User{}

// 	if err := db.Where("nick = ? AND password = ?", loginUser.Nick, loginUser.Password).First(&user).Error; err == nil {
// 		returnData := make(map[string]interface{})
// 		returnData["userInfo"] = user
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "验证成功", Data: returnData}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "用户名或密码错误", Data: false}
// 	}
// }

// func (model *User) GetAllUser(offset int, limit int, nick string, email string) helper.ReturnType {
// 	var users []User
// 	where := "nick like ? AND mail like ?"
// 	var count int

// 	db.Model(&User{}).Where(where, "%"+nick+"%", "%"+email+"%").Count(&count)

// 	err := db.Offset(offset).
// 		Limit(limit).
// 		Where(where, "%"+nick+"%", "%"+email+"%").
// 		Find(&users).
// 		Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
// 			Data: map[string]interface{}{
// 				"users": users,
// 				"count": count,
// 			},
// 		}
// 	}
// }

// func (model *User) GetUserByID(userID int) helper.ReturnType {
// 	var getUser User

// 	err := db.Select([]string{"user_id", "nick", "realname", "school", "major", "class", "mail"}).Where("user_id = ?", userID).First(&getUser).Error
// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: getUser}
// 	}
// }

// //GetUserByNick 由用户名获取用户
// func (model *User) GetUserByNick(userNick string) helper.ReturnType {
// 	var getUser User

// 	err := db.Select([]string{"user_id", "mail"}).Where("nick = ?", userNick).First(&getUser).Error
// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	}
// 	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: getUser}
// }