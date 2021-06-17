package com.liux.blog.controller.blog

import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.STATE_COMMENT_ENABLE
import com.liux.blog.bean.vo.CommentVO
import com.liux.blog.getIp
import com.liux.blog.getUserAgent
import com.liux.blog.service.ArticleService
import com.liux.blog.service.CommentService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.*
import javax.servlet.http.HttpServletRequest

@RestController
@RequestMapping("/article/{id}/comment")
class ArticleCommentController {

    @Autowired
    private lateinit var articleService: ArticleService
    @Autowired
    private lateinit var commentService: CommentService

    @GetMapping
    fun query(
        @PathVariable("id") articleId: Int,
        @RequestParam("subjectId") subjectId: Int,
        @RequestParam("page") page: Int,
    ): Resp<*> {
        val comments = commentService.listByArticle(articleId, page)
        val commentVO = CommentVO.of(comments)
        return Resp.succeed(commentVO)
    }

    @PutMapping
    fun insert(
        request: HttpServletRequest,
        @PathVariable("id") articleId: Int,
        @RequestParam("subjectId") subjectId: Int,
        @RequestParam("parentId") parentId: Int?,
        @RequestParam("nickname") nickname: String,
        @RequestParam("email") email: String,
        @RequestParam("link") link: String?,
        @RequestParam("content") content: String,
    ): Resp<*> {
        val ip = request.getIp()
        val ua = request.getUserAgent()

        if (nickname.isEmpty()) {
            return Resp.failed("昵称不能为空")
        }
        if (email.isEmpty()) {
            return Resp.failed("邮箱不能为空")
        }
        var localLink = link
        if (localLink.isNullOrEmpty()) localLink = null
        if (content.isEmpty()) {
            return Resp.failed("内容不能为空")
        }

        val article = articleService.getArticleById(articleId) ?: return Resp.failed("文章不存在")

        if (article.enableComment != STATE_COMMENT_ENABLE) {
            return Resp.failed("文章禁止评论")
        }

        if (parentId != null) {
            commentService.getCommentById(parentId) ?: return Resp.failed("父评论不存在")
        }

        commentService.addByBlog(articleId, parentId, nickname, email, localLink, content, ip, ua)
        return Resp.succeed("提交成功，请等待审核。")
    }
}
