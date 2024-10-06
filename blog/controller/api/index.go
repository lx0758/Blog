package api

import (
	"blog/controller"
	"blog/service"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	stringadapter "github.com/casbin/casbin/v2/persist/string-adapter"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IndexController struct {
	controller.RestController

	sessionService *service.SessionService
}

func (c *IndexController) OnInitController() {
	c.sessionService = service.GetService[*service.SessionService](c.sessionService)

	c.Group.Use(c.Recovery)
	c.Group.Use(c.Authorize)

	c.AddController("article", &ArticleController{})
	c.AddController("captcha", &CaptchaController{})
	c.AddController("category", &CategoryController{})
	c.AddController("comment", &CommentController{})
	c.AddController("config", &ConfigController{})
	c.AddController("dashboard", &DashboardController{})
	c.AddController("features", &FeatureController{})
	c.AddController("file", &FileController{})
	c.AddController("fragment", &FragmentController{})
	c.AddController("link", &LinkController{})
	c.AddController("session", &SessionController{})
	c.AddController("tag", &TagController{})
	c.AddController("url", &UrlController{})
	c.AddController("user", &UserController{})
}

func (c *IndexController) Recovery(context *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.RestFailed(context, controller.API_STATUS_ERROR, err.(error).Error())
		}
	}()
	context.Next()
}

var cModel, _ = model.NewModelFromString(`
[request_definition]
r = sub, obj, act
[policy_definition]
p = sub, obj, act
[role_definition]
g = _, _
[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && keyMatch(r.act, p.act)
[policy_effect]
e = some(where (p.eft == allow))
`)
var cAdapter = stringadapter.NewAdapter(`
g, -1, gGuest
p, gGuest, /api/session, POST
p, gGuest, /api/captcha, GET

g, 1, gAdmin
p, gAdmin, /api/*, *
`)
var cEnforcer, _ = casbin.NewEnforcer(cModel, cAdapter)

const (
	KEY_USER_ID = "userId"
	KEY_USER    = "user"
)

func (c *IndexController) Authorize(context *gin.Context) {
	userId := -1
	user := c.sessionService.GetLoggedUser(context)
	if user != nil {
		userId = user.Id
	}
	allow, _ := cEnforcer.Enforce(
		strconv.Itoa(userId),
		context.FullPath(),
		context.Request.Method,
	)
	if allow {
		context.Set(KEY_USER_ID, userId)
		context.Set(KEY_USER, user)
		context.Next()
	} else {
		c.RestFailed(context, controller.API_STATUS_UNAUTHORIZED, "无访问权限")
	}
}
