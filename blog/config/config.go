package config

import (
	bloglogger "blog/logger"
	"blog/res"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var gConf *Conf

func Config() *Conf {
	if gConf == nil {
		content, err := os.ReadFile("application.yaml")
		if err != nil {
			fmt.Printf("Read applocation.yaml error: %s\n", err)
		}

		conf := &Conf{}
		err = yaml.Unmarshal(content, &conf)
		if err != nil {
			fmt.Printf("Read applocation.yaml error: %s\n", err)
		}
		gConf = conf
	}
	return gConf
}

func Init(engine *gin.Engine) {
	if Config().Server.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.LoggerWithConfig(gin.LoggerConfig{
		Output: bloglogger.Writer(),
	})

	// 使用 engine.SetHTMLTemplate 自定义模板引擎时, 通过 engine.SetFuncMap 设置自定义函数无效
	funcMap := template.FuncMap{
		"filterChinese": func(param string) (string, error) {
			re := regexp.MustCompile("[\u4e00-\u9fa5]")
			return re.ReplaceAllString(param, ""), nil
		},
		"nowYear": func() (string, error) {
			return strconv.Itoa(time.Now().Year()), nil
		},
	}
	templateInstance, _ := template.New("").Funcs(funcMap).ParseFS(res.TemplatesFS, "html/*.html")
	engine.SetHTMLTemplate(templateInstance)

	engine.StaticFS("admin", http.FS(res.AdminStaticFS))
	engine.StaticFS("blog", http.FS(res.BlogStaticFS))

	storeKey := Config().Session.StoreKey
	if storeKey == "" {
		storeKey = strconv.FormatInt(time.Now().UnixMilli(), 10)
	}
	store := cookie.NewStore([]byte(storeKey))
	engine.Use(sessions.Sessions("session", store))

	engine.Use(cors.Default())
}

type Conf struct {
	Server     Server     `yaml:"server"`
	DataSource DataSource `yaml:"datasource"`
	Session    Session    `yaml:"session"`
}

type Server struct {
	Port    int  `yaml:"port"`
	Release bool `yaml:"release"`
}

type DataSource struct {
	Debug bool   `yaml:"debug"`
	Type  string `yaml:"type"`
	Dsn   string `yaml:"dsn"`
}

type Session struct {
	StoreKey string `yaml:"store_key"`
}
