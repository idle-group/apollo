package admin

import (
	"net/http"

	"github.com/chalvern/apollo/app/model"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// CommentsList 评论列表
func CommentsList(c *gin.Context) {
	c.Set(PageTitle, "评论列表")
	page := service.QueryPage(c)
	comments, allPage, err := service.CommentsQueryWithContext(c, true, "id desc")

	if err != nil {
		sugar.Errorf("CommentsList-获取 Comments 出错:%s", err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Timeout": 3,
		})
		return
	}

	html(c, http.StatusOK, "admin/comments/list.tpl", gin.H{
		"Comments":    comments,
		"CurrentPage": page,
		"TotalPage":   allPage,
	})
}

// TODO VaildStu 验证用户实名身份

func VaildStu(c *gin.Context) {
	c.Set(PageTitle, "用户实名审核")
	page := service.QueryPage(c)
	users, allPage, err := service.UsersQueryWithContext(c)

	if err != nil {
		sugar.Errorf("HomeIndex-获取 Shares 出错:%s", err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Timeout": 3,
		})
		return
	}

	html(c, http.StatusOK, "admin/users/vaild_stu.tpl", gin.H{
		"Users":       users,
		"CurrentPage": page,
		"TotalPage":   allPage,
	})
}

func ApproveStu(c *gin.Context) {
	print("asdadad")
	uidString := c.Query("uid")
	user, _ := service.UserFindByUID(uidString)
	c.Set(PageTitle, user.Nickname)

	form := struct {
		Priority int `form:"priority" binding:"required"`
	}{}
	c.ShouldBind(&form)

	userNew := model.User{
		Priority: form.Priority,
		StuVaild: true,
	}
	userNew.ID = user.ID
	err := service.UserUpdates(&userNew)
	if err != nil {
		sugar.Errorf("更新失败 %v", err)
	}
	html(c, http.StatusOK, "notify/success.tpl", gin.H{
		"Info":         "已更新",
		"Timeout":      3,
		"RedirectURL":  "/admin/account/list",
		"RedirectName": "用户列表",
	})
}
