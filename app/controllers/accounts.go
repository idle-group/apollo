package controllers

import (
	"fmt"
	"net/http"

	// "github.com/chalvern/apollo/app/mailer"
	"github.com/chalvern/apollo/app/service"
	"github.com/chalvern/apollo/configs/initializer"
	"github.com/chalvern/apollo/tools/jwt"
	"github.com/chalvern/sugar"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// SigninGet 获取登录页面
func SigninGet(c *gin.Context) {
	c.Set(PageTitle, "登陆")
	htmlOfOk(c, "account/signin.tpl", gin.H{})
}

// SignInPost 登陆
func SignInPost(c *gin.Context) {
	c.Set(PageTitle, "登陆")
	form := struct {
		Email     string `form:"email" binding:"required,email,lenlte=50"`
		Password  string `form:"password" binding:"required,lengte=8"`
		CaptchaID string `form:"captcha_id" binding:"required"`
		Captcha   string `form:"captcha" binding:"required"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("SigninPost Bind form Error: %s", errs.Error())
		html(c, http.StatusOK, "account/signin.tpl", gin.H{
			FlashError: "请检查邮箱、密码、验证码内容及格式是否填写正确",
		})
		return
	}

	// 验证码校验
	if !initializer.Captcha.Verify(form.CaptchaID, form.Captcha) {
		html(c, http.StatusBadRequest, "account/signin.tpl", gin.H{
			FlashError: "验证码错误",
		})
		return
	}

	u, err := service.UserSigninByEmail(form.Email, form.Password)
	if err != nil {
		sugar.Warnf("邮箱 %s 登录失败，密码错误。 err: %v", form.Email, err)
		html(c, http.StatusBadRequest, "account/signin.tpl", gin.H{
			FlashError: "邮箱未注册或密码错误",
		})
		return
	}

	// 设置 cookie
	token, err := jwt.NewToken(map[string]interface{}{
		"email": u.Email,
	})
	if err != nil {
		sugar.Errorf("SigninPost-NewToken-err: %s", err.Error())
		return
	}
	setJustCookie(c, token)

	c.Set("user", u)
	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         "登陆成功 😆😆😆",
		"Timeout":      3,
		"RedirectURL":  "/",
		"RedirectName": "主页",
	})
}

// SignupGet 获取注册页面
func SignupGet(c *gin.Context) {
	c.Set(PageTitle, "注册")
	html(c, http.StatusOK, "account/signup.tpl", gin.H{})
}

// SignUpPost 注册
func SignUpPost(c *gin.Context) {
	c.Set(PageTitle, "注册")
	form := struct {
		Email     string `form:"email" binding:"required,email,lenlte=100"`
		Password  string `form:"password" binding:"required,lengte=8,lenlte=128"`
		NickName  string `form:"nick_name" binding:"required,lengte=1,lenlte=50"`
		Password2 string `form:"password2" binding:"required,gtefield=Password,ltefield=Password"`
		StuId     string `form:"stu_id" binding:"required"`
		StuName	string `form:"stu_name" binding:"required"`
		CaptchaID string `form:"captcha_id" binding:"required"`
		Captcha   string `form:"captcha" binding:"required"`
	}{}

	if errs := c.ShouldBind(&form); errs != nil {
		sugar.Warnf("SigninPost Bind form Error: %s", errs.Error())
		// errors := errs.(validator.ValidationErrors)
		html(c, http.StatusOK, "account/signup.tpl", gin.H{
			FlashError: "请检查邮箱、密码、验证码内容及格式是否填写正确",
		})
		return
	}

	// 验证码校验
	if !initializer.Captcha.Verify(form.CaptchaID, form.Captcha) {
		html(c, http.StatusBadRequest, "account/signup.tpl", gin.H{
			FlashError: "验证码错误",
		})
		return
	}

	// 在这里进行 文件的处理
	file, _ := c.FormFile("upload_image")
	dst := "./upload_images/" + uuid.NewV4().String() + file.Filename
	c.SaveUploadedFile(file, dst)

	newUser, err := service.UserSignup(form.Email, form.Password, form.NickName, form.StuId, form.StuName,dst)
	if err != nil {
		html(c, http.StatusBadRequest, "account/signup.tpl", gin.H{
			FlashError: "创建用户失败，邮箱已注册",
		})
		return
	}
	fmt.Println(newUser)

	// // 发送验证邮件
	// err = mailer.AccountValidEmail(form.Email, form.NickName, newUser.EmailValidToken)
	// if err != nil {
	// 	sugar.Warnf("发送验证邮件失败，email: %s", form.Email)
	// }

	htmlOfOk(c, "notify/success.tpl", gin.H{
		"Info":         fmt.Errorf("注册成功 😆😆😆，(等待管理员验证)"),
		"Timeout":      3,
		"RedirectURL":  "/signin",
		"RedirectName": "登陆页",
	})

}

// SignOut 注销登陆
func SignOut(c *gin.Context) {
	c.Set(PageTitle, "注销")
	c.Set("user", nil)
	expireCookie(c)
	html(c, http.StatusOK, "notify/success.tpl", gin.H{
		"Info":         "已注销",
		"Timeout":      3,
		"RedirectURL":  "/",
		"RedirectName": "首页",
	})
}

// AccountValidEmailHandler 验证邮箱
func AccountValidEmailHandler(c *gin.Context) {
	c.Set(PageTitle, "邮箱验证")
	mail := c.Query("mail")
	token := c.Query("token")
	if mail == "" || token == "" {
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "参数无效",
		})
		return
	}
	err := service.UserValidEmail(mail, token)
	if err != nil {
		sugar.Warnf("用户校验邮箱出错：%s", err.Error)
		html(c, http.StatusOK, "notify/error.tpl", gin.H{
			"FlashError": "邮箱未注册或 token 已过期",
		})
		return
	}

	html(c, http.StatusOK, "notify/success.tpl", gin.H{
		"Info":         "验证成功 😆😆😆",
		"Timeout":      5,
		"RedirectURL":  "/signin",
		"RedirectName": "登陆页",
	})
}

