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
	_ "net/http/pprof"
	"regexp"
	"strconv"
	"strings"
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
	pprofConf := config.Config().Pprof
	if pprofConf.Enable {
		address := pprofConf.Address
		logger.Printf("Start pprof server at: http://%s", address)
		go func() {
			logger.Panicf(http.ListenAndServe(address, nil).Error())
		}()
	}

	serverConf := config.Config().Server
	var ginMode string
	if serverConf.Debug {
		ginMode = gin.DebugMode
	} else {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	engine := gin.New()
	engine.With(
		InstallLogger,
		InstallRecovery,
		InstallRemoterIPHeader,
		InstallTemplate,
		InstallSession,
		InstallStatic,
		InstallCors,
	)

	controller.AddController(engine, "", htmlIndexController)
	controller.AddController(engine, "api", apiIndexController)

	err := engine.Run(serverConf.Address)
	if err != nil {
		logger.Panicf("Run server error:%s\n", err)
	}
}

func InstallLogger(engine *gin.Engine) {
	serverConf := config.Config().Server
	if serverConf.Debug {
		engine.Use(gin.Logger())
		return
	}
	skipPaths := []string{
		"/blog/",
		"/admin/",
	}
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Skip: func(context *gin.Context) bool {
			path := context.Request.URL.Path
			for _, skipPath := range skipPaths {
				if strings.HasPrefix(path, skipPath) {
					return true
				}
			}
			return false
		},
	}))
}

func InstallRecovery(engine *gin.Engine) {
	engine.Use(func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 全局 panic 处理
				logger.Printf("Recovered from panic:%s\n", err)
				htmlIndexController.RenderError(context, err.(error))
			} else if context.Writer.Status() == http.StatusNotFound {
				// 处理 Gin 路由没有命中的场景
				notFoundError := errors.New("Not found " + context.Request.URL.Path)
				logger.Printf("Recovered from NotFound:%s\n", notFoundError)
				htmlIndexController.RenderError(context, notFoundError)
			}
		}()
		context.Next()
	})
}

func InstallRemoterIPHeader(engine *gin.Engine) {
	engine.RemoteIPHeaders = []string{
		"X-Forwarded-For",
		"Proxy-Client-IP",
		"WL-Proxy-Client-IP",
		"HTTP_CLIENT_IP",
		"X-Real-IP",
	}
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
		"navLink": func(index int, anchor string, title string) template.HTML {
			builder := strings.Builder{}
			builder.WriteString(fmt.Sprintf("<a class=\"nav-link\" href=\"#%s\">", anchor))
			builder.WriteString(fmt.Sprintf("    <span class=\"nav-number\">%d.</span>", index+1))
			builder.WriteString(fmt.Sprintf("    <span class=\"nav-text\">%s</span>", title))
			builder.WriteString(fmt.Sprintf("</a>"))
			return template.HTML(builder.String())
		},
	}
	templateInstance, err := template.New("").Funcs(funcMap).ParseFS(res.TemplatesFS, "html/*.gohtml")
	if err != nil {
		logger.Panicf("Failed initialize template:%s\n", err)
	}
	engine.SetHTMLTemplate(templateInstance)
}

func InstallSession(engine *gin.Engine) {
	sessionConf := config.Config().Session
	storeKey := sessionConf.StoreKey
	if storeKey == "" {
		storeKey = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	store := memstore.NewStore([]byte(storeKey))
	engine.Use(sessions.Sessions("session", store))
}

func InstallStatic(engine *gin.Engine) {
	engine.StaticFS("/blog", http.FS(res.BlogStaticFS))
	engine.StaticFS("/admin", http.FS(res.AdminStaticFS))
	engine.StaticFS("/files", gin.Dir("files", false))
}

func InstallCors(engine *gin.Engine) {
	engine.Use(cors.Default())
}
