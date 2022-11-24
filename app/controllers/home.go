package controllers

import (
	"net/http"

	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
)

// HomeIndex Home 首页
func HomeIndex(c *gin.Context) {
	c.Set(PageTitle, "hfut校园bbs（分享日常、品鉴生活）")
	tabString := c.Query("t")
	if tabString == "" {
		tabString = "0"
	}

	page := service.QueryPage(c)
	shares, allPage, err := service.SharesQueryWithContext(c, true)
	if err != nil {
		sugar.Errorf("HomeIndex-获取 Shares 出错:%s", err.Error())
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"Timeout": 3,
		})
		return
	}

	html(c, http.StatusOK, "home/index.tpl", gin.H{
		"Shares":      shares,
		"TabIndex":    tabString,
		"CurrentPage": page,
		"TotalPage":   allPage,
		"SideTags":    service.TagsRecommendQuery(),
	})
}

// HomeAboutHandler 关于
func HomeAboutHandler(c *gin.Context) {
	c.Set(PageTitle, "关于")
	html(c, http.StatusOK, "home/about.tpl", gin.H{})
}
