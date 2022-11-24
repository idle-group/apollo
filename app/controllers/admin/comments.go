package admin

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

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
	users, allPage, err := service.UsersQueryWithContext(c, "stu_vaild=0")
	if err != nil {
		sugar.Errorf("HomeIndex-获取 Shares 出错:%s", err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Timeout": 3,
		})
		return
	}

	for i := 0; i < len(users); i++ {
		photopath := users[i].StuPhoto
		file, err := os.Open(photopath)
		if err != nil {
			sugar.Errorf("获取 用户图片 出错:%s", err.Error())
			html(c, http.StatusOK, "notify/error.tpl", gin.H{
				"Timeout": 3,
			})
			return
		}
		defer file.Close()
		imgByte, _ := ioutil.ReadAll(file)
		mimeType := http.DetectContentType(imgByte)
		switch mimeType {
		case "image/jpeg":
			fmt.Println("jpeg")
			users[i].StuPhotoBase64 = template.URL("data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte))
		case "image/png":
			fmt.Println("png")
			users[i].StuPhotoBase64 = template.URL("data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte))
			fmt.Println(users[i].StuPhoto)
		}
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
		Priority: 1,
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
