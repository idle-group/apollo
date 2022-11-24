package admin

import (
	"net/http"

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

// VaildStu 验证用户实名身份

func VaildStu(c *gin.Context) {
	c.Set(PageTitle,"用户实名审核")
	html(c,http.StatusOK,"admin/users/vaild_stu.tpl",gin.H{})
}