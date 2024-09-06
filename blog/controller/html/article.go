package html

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
	controller.Controller
}

func (c *ArticleController) OnInitController() {
	c.Group.GET(":key", c.getArticle)
	c.Group.GET(":key/comment", c.getArticleComment)
	c.Group.POST(":key/comment", c.postArticleComment)
}

func (c *ArticleController) getArticle(context *gin.Context) {
	_ = context.Param("key")
	c.Render(context, http.StatusOK, "article.html", nil)
}

func (c *ArticleController) getArticleComment(context *gin.Context) {

}

func (c *ArticleController) postArticleComment(context *gin.Context) {

}
