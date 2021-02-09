package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/panel/model"
	"container/list"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type menuItem struct {
	Title 	string `json:"title"`
	Icon	string `json:"icon"`
	Href	string `json:"href"`
	Target 	string `json:"target"`
	Child 	[]menuItem `json:"child"`
}


func haveAuth(c *gin.Context, authQuery string) int {
	session := sessions.Default(c)
	id := session.Get("user_id")
	if  id == nil {
		return common.UnLoggedIn
	} else if session.Get("identity").(uint) == 0 {
		return common.UnAuthed
	}
	_, auths, err := getUserAllAuth(id.(int))
	if err != nil {
		return common.AuthError
	} else {
		for _, auth := range auths {
			if auth == authQuery {
				return common.Authed
			}
		}
		return common.UnAuthed
	}
}

func getUserAllAuth(userID int) ([]menuItem, []string, error) {
	authModel := model.Auth{}

	if res := authModel.GetUserAllAuth(userID); res.Status == common.CodeSuccess {
		auths := res.Data.([]model.Auth)
		authsLeft := map[int]model.Auth{}
		var authName []string
		var menu []menuItem

		childMenu := map[int][]menuItem{}
		hasChild := map[int]int{}
		menuItemCount := 0

		for _, auth := range auths {
			if auth.Type == 2 {
				authName = append(authName, auth.Title)
			} else {
				hasChild[auth.Aid] = 0
				childMenu[auth.Aid] = make([]menuItem, 0)
				authsLeft[auth.Aid] = auth
				if auth.Type == 0 {
					menuItemCount++
				}
			}
		}

		for _, auth := range authsLeft {
			hasChild[auth.Parent]++
		}
		queue := list.New()
		for _, auth := range authsLeft {
			if hasChild[auth.Aid] == 0 {
				queue.PushBack(auth)
			}
		}

		for queue.Len() > 0 {
			auth := queue.Front().Value.(model.Auth)
			queue.Remove(queue.Front())
			item := menuItem{
				Title: auth.Title,
				Icon: auth.Icon,
				Href: auth.Href,
				Target: auth.Target,
				Child: childMenu[auth.Aid],
			}
			childMenu[auth.Parent] = append(childMenu[auth.Parent], item)
			hasChild[auth.Parent]--
			if hasChild[auth.Parent] == 0 {
				parentAuth := authsLeft[auth.Parent]
				if parentAuth.Type == 0 {
					menu = append(menu, menuItem{
						Title: parentAuth.Title,
						Icon: parentAuth.Icon,
						Href: parentAuth.Href,
						Target: parentAuth.Target,
						Child: childMenu[parentAuth.Aid],
					})
				} else {
					queue.PushBack(authsLeft[auth.Parent])
				}
			}
		}

		return menu, authName, nil
	} else {
		return nil, nil, errors.New("获取权限列表错误")
	}

}