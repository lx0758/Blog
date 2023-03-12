package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.CommentStatusEnum
import com.liux.blog.bean.po.isValid
import com.liux.blog.bean.vo.api.CommentVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.getIp
import com.liux.blog.getUserAgent
import com.liux.blog.service.ArticleService
import com.liux.blog.service.CommentService
import com.liux.blog.service.UserService
import jakarta.servlet.http.HttpServletRequest
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/comment")
class CommentController {
    
    @Autowired
    private lateinit var commentService: CommentService
    @Autowired
    private lateinit var articleService: ArticleService
    @Autowired
    private lateinit var userService: UserService

    @GetMapping
    fun query(
        @RequestParam("article", required = false) article: Int?,
        @RequestParam("author", required = false) author: String?,
        @RequestParam("email", required = false) email: String?,
        @RequestParam("ip", required = false) ip: String?,
        @RequestParam("content", required = false) content: String?,
        @RequestParam("status", required = false) status: Int?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
        @RequestParam("orderName", required = false) orderName: String?,
        @RequestParam("orderMethod", required = false) orderMethod: String?,
    ): Resp<PaginationListVO<CommentVO>> {
        val commentPage = commentService.listByAdmin(article, author, email, ip, content, status, pageNum, pageSize, orderName, orderMethod)
        val comments = commentPage.map { CommentVO.of(it) }
        return Resp.succeed(PaginationListVO.of(comments, commentPage))
    }

    @PostMapping
    fun add(
        request: HttpServletRequest,
        @CurrentUserId userId: Int,
        @RequestParam("articleId") articleId: Int,
        @RequestParam("parentId") parentId: Int,
        @RequestParam("content") content: String,
        @RequestParam("emailNotify") emailNotify: Boolean,
    ): Resp<*> {
        articleService.getByAdmin(articleId) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "文章不存在")
        commentService.getCommentById(parentId) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "父评论不存在")
        if (content.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "内容不能为空")
        }
        val user = userService.getById(userId)!!
        val ip = request.getIp()
        val ua = request.getUserAgent()
        commentService.addByAdmin(userId, user.nickname!!, user.email!!, articleId, parentId, content, ip, ua, emailNotify)
        return Resp.succeed()
    }

    @PutMapping("{id}")
    fun update(
        @PathVariable("id") id: Int,
        @RequestParam("status") status: Int,
    ): Resp<*> {
        if (!CommentStatusEnum.values().isValid(status)) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "状态不正确")
        }
        val updateRow = commentService.updateByAdmin(id, status)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val deleteRow = commentService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }
}