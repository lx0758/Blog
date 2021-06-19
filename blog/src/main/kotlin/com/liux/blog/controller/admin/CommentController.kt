package com.liux.blog.controller.admin

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.CommentItemVO
import com.liux.blog.bean.vo.PaginationListVO
import com.liux.blog.service.CommentService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/comment")
class CommentController {
    
    @Autowired
    private lateinit var commentService: CommentService

    @GetMapping
    fun query(@RequestParam("pageNum") pageNum: Int, @RequestParam("pageSize") pageSize: Int): Resp<PaginationListVO<CommentItemVO>> {
        val commentPage = commentService.listByPage(pageNum, pageSize)
        val comments = commentPage.map { CommentItemVO.of(it) }
        return Resp.succeed(PaginationListVO.of(comments, commentPage))
    }
}