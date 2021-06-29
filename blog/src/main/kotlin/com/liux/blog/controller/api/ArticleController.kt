package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.ArticleItemVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.service.ArticleService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/api/article")
class ArticleController {

    @Autowired
    private lateinit var articleService: ArticleService

    @GetMapping
    fun query(
        @RequestParam("title", required = false) title: String?,
        @RequestParam("category", required = false) category: Int?,
        @RequestParam("format", required = false) format: Int?,
        @RequestParam("comment", required = false) comment: Int?,
        @RequestParam("status", required = false) status: Int?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
    ): Resp<PaginationListVO<ArticleItemVO>> {
        val articlePage = articleService.listByAdmin(title, category, format, comment, status, pageNum, pageSize)
        val articles = articlePage.map { ArticleItemVO.of(it) }
        return Resp.succeed(PaginationListVO.of(articles, articlePage))
    }

    @PostMapping
    fun insert(): Resp<*> {
        TODO()
    }

    @PutMapping("{id}")
    fun update(@PathVariable("id") id: Int): Resp<*> {
        TODO()
    }

    @PutMapping("{id}/status")
    fun updateStatus(@PathVariable("id") id: Int, @RequestParam("status") status: Int): Resp<*> {
        TODO()
    }

    @DeleteMapping("{id}")
    fun delete(@PathVariable("id") id: Int): Resp<*> {
        TODO()
    }
}