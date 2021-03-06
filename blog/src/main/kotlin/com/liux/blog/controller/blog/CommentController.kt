package com.liux.blog.controller.blog

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.CommentVO
import com.liux.blog.service.CommentService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.*
import javax.servlet.http.HttpServletRequest

@RestController
@RequestMapping("/comment")
class CommentController {

    @Autowired
    private lateinit var commentService: CommentService

    @GetMapping
    fun query(
        @RequestParam("subjectId") articleId: Int,
        @RequestParam("page") page: Int,
    ): Resp<*> {
        val comments = commentService.listByArticle(articleId, page)
        val commentVO = CommentVO.of(comments)
        return Resp.succeed(commentVO)
    }

    @PutMapping
    fun insert(
        request: HttpServletRequest,
        @RequestParam("subjectId") articleId: Int,
        @RequestParam("parentId") parentId: Int?,
        @RequestParam("nickname") nickname: String,
        @RequestParam("email") email: String,
        @RequestParam("link") link: String?,
        @RequestParam("content") content: String,
    ): Resp<*> {
        val headers = arrayOf(
            "X-Forwarded-For",
            "Proxy-Client-IP",
            "WL-Proxy-Client-IP",
            "HTTP_CLIENT_IP",
            "X-Real-IP",
        )
        var index = 0
        var ip: String? = null
        while (ip.isNullOrEmpty() && index < headers.size) {
            ip = request.getHeader(headers[index])
            index ++
        }
        if (ip.isNullOrEmpty()) {
            ip = request.remoteAddr
        }

        val ua = request.getHeader("User-Agent")

        if (nickname.isEmpty()) {
            return Resp.error("昵称不能为空")
        }
        if (email.isEmpty()) {
            return Resp.error("邮箱不能为空")
        }
        var localLink = link
        if (localLink.isNullOrEmpty()) localLink = null
        if (content.isEmpty()) {
            return Resp.error("内容不能为空")
        }

        commentService.addByBlog(articleId, parentId, nickname, email, localLink, content, ip, ua)
        return Resp.succeed()
    }
}
