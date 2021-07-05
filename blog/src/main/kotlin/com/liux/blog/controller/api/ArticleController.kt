package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.ArticleStatusEnum
import com.liux.blog.bean.po.isValid
import com.liux.blog.bean.vo.api.ArticleItemVO
import com.liux.blog.bean.vo.api.ArticleVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.service.ArticleService
import com.liux.blog.service.CategoryService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/article")
class ArticleController {

    @Autowired
    private lateinit var articleService: ArticleService
    @Autowired
    private lateinit var categoryService: CategoryService

    @GetMapping
    fun query(
        @RequestParam("title", required = false) title: String?,
        @RequestParam("category", required = false) category: Int?,
        @RequestParam("enableComment", required = false) enableComment: Boolean?,
        @RequestParam("status", required = false) status: Int?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
    ): Resp<PaginationListVO<ArticleItemVO>> {
        val articlePage = articleService.listByAdmin(title, category, enableComment, status, pageNum, pageSize)
        val articles = articlePage.map { ArticleItemVO.of(it) }
        return Resp.succeed(PaginationListVO.of(articles, articlePage))
    }

    @GetMapping("{id}")
    fun queryInfo(
        @PathVariable("id") id: Int,
    ): Resp<ArticleVO> {
        val article = articleService.getByAdmin(id) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "文章不存在")
        val articleVO = ArticleVO.of(article)
        return Resp.succeed(articleVO)
    }

    @PostMapping
    fun add(
        @CurrentUserId userId: Int,
        @RequestParam("title") title: String,
        @RequestParam("categoryId") categoryId: Int,
        @RequestParam("content") content: String,
        @RequestParam("url", required = false) url: String?,
        @RequestParam("weight", required = false) weight: Int?,
        @RequestParam("enableComment", required = false) enableComment: Boolean,
        @RequestParam("status") status: Int,
        @RequestParam("tags", required = false) tags: Array<String>?,
    ): Resp<ArticleVO> {
        checkParams(title, categoryId, content, url, weight, status)
        val article = articleService.addByAdmin(
            userId,
            title,
            categoryId,
            content,
            url,
            weight,
            enableComment,
            status,
            tags,
        )
        val articleVO = ArticleVO.of(article)
        return Resp.succeed(articleVO)
    }

    @PutMapping("{id}")
    fun update(
        @PathVariable("id") id: Int,
        @RequestParam("title") title: String,
        @RequestParam("categoryId") categoryId: Int,
        @RequestParam("content") content: String,
        @RequestParam("url", required = false) url: String?,
        @RequestParam("weight", required = false) weight: Int?,
        @RequestParam("enableComment", required = false) enableComment: Boolean,
        @RequestParam("status") status: Int,
        @RequestParam("tags", required = false) tags: Array<String>?,
    ): Resp<*> {
        checkParams(title, categoryId, content, url, weight, status)
        val updateRow = articleService.updateByAdmin(
            id,
            title,
            categoryId,
            content,
            url,
            weight,
            enableComment,
            status,
            tags,
        )
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @PutMapping("{id}/status")
    fun updateStatus(
        @PathVariable("id") id: Int,
        @RequestParam("status") status: Int,
    ): Resp<*> {
        val updateRow = articleService.updateStatusByAdmin(id, status)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int
    ): Resp<*> {
        val deleteRow = articleService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }

    private fun checkParams(
        title: String,
        categoryId: Int,
        content: String,
        url: String?,
        weight: Int?,
        status: Int
    ) {
        if (title.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "标题不能为空")
        }
        if (categoryService.getById(categoryId) == null) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "分类不存在")
        }
        if (content.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "内容不能为空")
        }
        if (!ArticleStatusEnum.values().isValid(status)) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "状态类型不正确")
        }
    }
}