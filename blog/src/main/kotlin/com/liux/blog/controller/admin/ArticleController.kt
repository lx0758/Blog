package com.liux.blog.controller.admin

import com.liux.blog.bean.vo.PaginationListVO
import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.ArticleItemVO
import com.liux.blog.service.ArticleService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/api/article")
class ArticleController {

    @Autowired
    private lateinit var articleService: ArticleService

    @GetMapping
    fun query(@RequestParam("pageNum") pageNum: Int, @RequestParam("pageSize") pageSize: Int): Resp<PaginationListVO<ArticleItemVO>> {
        val articlePage = articleService.listByPage(pageNum, pageSize)
        val articles = articlePage.map { ArticleItemVO.of(it) }
        return Resp.succeed(PaginationListVO.of(articles, articlePage))
    }

    @PutMapping
    fun add(): Resp<*> {
        TODO()
    }

    @PostMapping("{id}")
    fun update(@PathVariable("id") id: Int): Resp<*> {
        TODO()
    }

    @PostMapping("{id}/status")
    fun updateStatus(@PathVariable("id") id: Int, @RequestParam("status") status: Int): Resp<*> {
        TODO()
    }

    @DeleteMapping("{id}")
    fun delete(@PathVariable("id") id: Int): Resp<*> {
        TODO()
    }
}