package router

import (
	"github.com/chalvern/apollo/app/controllers/admin"
	i "github.com/chalvern/apollo/app/interceptors"
	"github.com/chalvern/apollo/app/model"
)

func adminRouterInit() {
	get("admin_home_page", "/admin",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.HomeIndex)
	get("admin_account_list", "/admin/account/list",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.AccountsList)
	get("admin_account_edit_get", "/admin/account/edit",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.AccountsEditGet)
	post("admin_account_edit_post", "/admin/account/edit",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.AccountsEditPost)

	get("admin_comments_list", "/admin/comments/list",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.CommentsList)
	get("admin_valid_stu", "/admin/account/vaild_stu", i.UserPriorityMiddleware(model.UserPrioritySuper), admin.VaildStu)

	post("admin_valid_stu_post", "/admin/account/vaild_stu/approve",
		i.UserPriorityMiddleware(model.UserPrioritySuper), admin.ApproveStu)
}
