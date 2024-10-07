package main

import (
	"blog/config"
	"blog/controller"
	"blog/controller/api"
	"blog/controller/html"
	"blog/logger"
	"blog/res"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var htmlIndexController = &html.IndexController{}
var apiIndexController = &api.IndexController{}

// go install github.com/swaggo/swag/cmd/swag@latest
// https://github.com/swaggo/swag/blob/master/operation.go
// swag init
//
// @title			个人博客
// @version		1.0
// @description	个人博客接口
// @contact.name	6x
// @contact.url	http://github.com/lx0758/
// @contact.email	lx0758@qq.com
// @host			127.0.0.1:8080
// @basePath		/
func main() {
	conf := config.Config().Server
	if conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	engine.With(
		InstallRecovery,
		InstallCors,
		InstallStatic,
		InstallTemplate,
		InstallSession,
	)

	controller.AddController(engine, "", htmlIndexController)
	controller.AddController(engine, "api", apiIndexController)

	port := conf.Port
	err := engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Panicf("Run server error:%s\n", err)
	}
}

func InstallRecovery(engine *gin.Engine) {
	engine.Use(func(context *gin.Context) {
		defer func() {
			if context.Writer.Status() == http.StatusNotFound {
				htmlIndexController.RenderError(context, errors.New("not found"))
			} else if err := recover(); err != nil {
				logger.Printf("Recovered from panic:%s\n", err)
				htmlIndexController.RenderError(context, err.(error))
			}
		}()
		context.Next()
	})
}

func InstallCors(engine *gin.Engine) {
	engine.Use(cors.Default())
}

func InstallStatic(engine *gin.Engine) {
	engine.StaticFS("/blog", http.FS(res.BlogStaticFS))
	engine.StaticFS("/admin", http.FS(res.AdminStaticFS))
	engine.StaticFS("/files", gin.Dir("files", false))
}

func InstallTemplate(engine *gin.Engine) {
	// 使用 engine.SetHTMLTemplate 自定义模板引擎时, 通过 engine.SetFuncMap 设置自定义函数无效
	funcMap := template.FuncMap{
		"filterChinese": func(param string) (string, error) {
			re := regexp.MustCompile("[\u4e00-\u9fa5]")
			return re.ReplaceAllString(param, ""), nil
		},
		"nowYear": func() (string, error) {
			return strconv.Itoa(time.Now().Year()), nil
		},
		"getYear": func(t time.Time) (int, error) {
			return t.Year(), nil
		},
		"formatDate": func(t time.Time) (string, error) {
			return t.Format(time.DateOnly), nil
		},
		"formatTime": func(t time.Time) (string, error) {
			return t.Format(time.TimeOnly), nil
		},
		"formatMMdd": func(t time.Time) (string, error) {
			return t.Format("01-02"), nil
		},
		"formatDateTime": func(t time.Time) (string, error) {
			return t.Format(time.DateTime), nil
		},
		"rawHtml": func(content string) template.HTML {
			return template.HTML(content)
		},
		"addInt": func(param1 int, param2 int) (int, error) {
			return param1 + param2, nil
		},
		"addFloat": func(param1 float32, param2 float32) (float32, error) {
			return param1 + param2, nil
		},
		"tagOpacity": func(count int) (float32, error) {
			return float32(count)/20 + 0.5, nil
		},
	}
	templateInstance, err := template.New("").Funcs(funcMap).ParseFS(res.TemplatesFS, "html/*.gohtml")
	if err != nil {
		logger.Panicf("Failed initialize template:%s\n", err)
	}
	engine.SetHTMLTemplate(templateInstance)
}

func InstallSession(engine *gin.Engine) {
	storeKey := config.Config().Session.StoreKey
	if storeKey == "" {
		storeKey = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	store := memstore.NewStore([]byte(storeKey))
	engine.Use(sessions.Sessions("session", store))
}
