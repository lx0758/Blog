package com.liux.blog.controller

import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.ArticleCommentStatusEnum
import com.liux.blog.bean.vo.CommentPageVO
import com.liux.blog.getIp
import com.liux.blog.getUserAgent
import com.liux.blog.service.ArticleService
import com.liux.blog.service.CaptchaService
import com.liux.blog.service.CommentService
import jakarta.servlet.http.HttpServletRequest
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/article/{id}/comment")
class ArticleCommentController {

    @Autowired
    private lateinit var articleService: ArticleService
    @Autowired
    private lateinit var commentService: CommentService
    @Autowired
    private lateinit var captchaService: CaptchaService

    @GetMapping
    fun query(
        @PathVariable("id") articleId: Int,
        @RequestParam("subjectId", required = false) subjectId: Int?,
        @RequestParam("page") page: Int,
    ): Resp<CommentPageVO> {
        val comments = commentService.listByArticle(articleId, page)
        val commentVO = CommentPageVO.of(comments)
        return Resp.succeed(commentVO)
    }

    @PostMapping
    fun add(
        request: HttpServletRequest,
        @RequestParam("captcha") captcha: String,
        @PathVariable("id") articleId: Int,
        @RequestParam("subjectId", required = false) subjectId: Int?,
        @RequestParam("parentId", required = false) parentId: Int?,
        @RequestParam("nickname") nickname: String,
        @RequestParam("email") email: String,
        @RequestParam("link", required = false) link: String?,
        @RequestParam("content") content: String,
    ): Resp<*> {
        if (captcha.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "验证码不能为空")
        }
        if (!captchaService.verify( request.session, captcha, CaptchaService.TYPE_COMMENT)) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "验证码错误")
        }
        if (nickname.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "昵称不能为空")
        }
        if (email.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "邮箱不能为空")
        }
        var localLink = link
        if (localLink.isNullOrEmpty()) localLink = null
        if (content.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "内容不能为空")
        }

        val article = articleService.getByBlog(articleId) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "文章不存在")

        if (article.commentStatus != ArticleCommentStatusEnum.ENABLE.value) {
            throw HttpClientErrorException(HttpStatus.FORBIDDEN, "文章禁止评论")
        }

        if (parentId != null) {
            commentService.getCommentById(parentId) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "父评论不存在")
        }

        val ip = request.getIp()
        val ua = request.getUserAgent()
        commentService.addByBlog(articleId, parentId, nickname, email, localLink, content, ip, ua)
        return Resp.succeed()
    }
}
