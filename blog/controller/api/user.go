package api

import (
	"blog/bean/api_vo"
	"blog/controller"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controller.RestController

	userService *service.UserService
}

func (c *UserController) OnInitController() {
	c.userService = service.GetService[*service.UserService](c.userService)

	c.Group.GET("", c.listUser)
	c.Group.GET("profile", c.queryProfile)
	c.Group.PUT("profile", c.updateProfile)
	c.Group.PUT("password", c.updatePassword)
}

type listUser struct {
	Username    *string `form:"username"`
	Nickname    *string `form:"nickname"`
	Status      *int    `form:"status"`
	PageNum     int     `form:"pageNum" binding:"required"`
	PageSize    int     `form:"pageSize" binding:"required"`
	OrderName   *string `form:"orderName"`
	OrderMethod *string `form:"orderMethod"`
}

// @Summary		list user
// @Description	list user
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		username		body	string	false	"username"
// @Param		nickname		body	string	false	"nickname"
// @Param		status			body	int	false	"status"
// @Param		pageNum			body	int		true	"pageNum"
// @Param		pageSize		body	int		true	"pageSize"
// @Param		orderName		body	string	false	"orderName"
// @Param		orderMethod		body	string	false	"orderMethod"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/user [GET]
func (c *UserController) listUser(context *gin.Context) {
	var listUser listUser
	if err := context.ShouldBind(&listUser); err != nil {
		c.RestValidationError(context, err)
	}

	users := c.userService.ListByAdmin(
		listUser.Username,
		listUser.Nickname,
		listUser.Status,
	)
	userVOs := make([]api_vo.UserVO, 0)
	for _, user := range users {
		userVO := api_vo.UserVO{}
		userVO.From(user)
		userVOs = append(userVOs, userVO)
	}
	c.RestSucceed(context, userVOs)
}

// @Summary		query user profile
// @Description	query user profile
// @Tags		user
// @Accept		json
// @Produce		json
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/user/profile [GET]
func (c *UserController) queryProfile(context *gin.Context) {
	userId := context.GetInt(KEY_USER_ID)
	user := c.userService.QueryUserById(userId)
	userVO := api_vo.UserVO{}
	userVO.From(*user)
	c.RestSucceed(context, userVO)
}

type updateProfile struct {
	Avatar      *string `form:"avatar"`
	Nickname    string  `form:"nickname" binding:"required"`
	Description *string `form:"description"`
	Email       string  `form:"email" binding:"required"`

	Github        *string `form:"github"`
	Weibo         *string `form:"weibo"`
	Google        *string `form:"google"`
	Twitter       *string `form:"twitter"`
	Facebook      *string `form:"facebook"`
	StackOverflow *string `form:"stackOverflow"`
	Youtube       *string `form:"youtube"`
	Instagram     *string `form:"instagram"`
	Skype         *string `form:"skype"`
}

// @Summary		update user profile
// @Description	update user profile
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		avatar			body	string	false	"avatar"
// @Param		nickname		body	string	true	"nickname"
// @Param		description		body	string	false	"description"
// @Param		email			body	string	true	"email"
// @Param		github			body	string	false	"github"
// @Param		weibo			body	string	false	"weibo"
// @Param		google			body	string	false	"google"
// @Param		twitter			body	string	false	"twitter"
// @Param		facebook		body	string	false	"facebook"
// @Param		stackOverflow	body	string	false	"stackOverflow"
// @Param		youtube			body	string	false	"youtube"
// @Param		instagram		body	string	false	"instagram"
// @Param		skype			body	string	false	"skype"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/user/profile [PUT]
func (c *UserController) updateProfile(context *gin.Context) {
	var updateProfile updateProfile
	if err := context.ShouldBind(&updateProfile); err != nil {
		c.RestValidationError(context, err)
	}

	userId := context.GetInt(KEY_USER_ID)
	result := c.userService.UpdateUserProfile(
		userId,
		updateProfile.Avatar,
		updateProfile.Nickname,
		updateProfile.Description,
		&updateProfile.Email,
		updateProfile.Github,
		updateProfile.Weibo,
		updateProfile.Google,
		updateProfile.Twitter,
		updateProfile.Facebook,
		updateProfile.StackOverflow,
		updateProfile.Youtube,
		updateProfile.Instagram,
		updateProfile.Skype,
	)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新资料失败")
	}
}

type updatePassword struct {
	OldPassword string `form:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" binding:"required"`
}

// @Summary		update user password
// @Description	update user password
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		oldPassword		body	string	true	"oldPassword"
// @Param		newPassword		body	string	true	"newPassword"
// @Success		200			{object}	string	"{"status": 0, "message": "", “data”: null}"
// @Failure		200			{object}	string	"{"status": 500, "message": "", “data”: null}"
// @Router		/api/user/password [PUT]
func (c *UserController) updatePassword(context *gin.Context) {
	var updatePassword updatePassword
	if err := context.ShouldBind(&updatePassword); err != nil {
		c.RestValidationError(context, err)
	}

	userId := context.GetInt(KEY_USER_ID)
	result := c.userService.UpdateUserPassword(
		userId,
		updatePassword.OldPassword,
		updatePassword.NewPassword,
	)
	if result {
		c.RestSucceed(context, nil)
	} else {
		c.RestError(context, "更新密码失败")
	}
}
