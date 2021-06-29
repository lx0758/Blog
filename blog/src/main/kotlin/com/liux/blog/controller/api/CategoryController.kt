package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.CategoryVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.service.CategoryService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/category")
class CategoryController {

    @Autowired
    private lateinit var categoryService: CategoryService

    @GetMapping
    fun query(
        @RequestParam("name", required = false) name: String?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
    ): Resp<PaginationListVO<CategoryVO>> {
        val categorys = categoryService.listByAdmin(name, pageNum, pageSize)
        val categoryVOs = categorys.map { CategoryVO.of(it) }
        return Resp.succeed(PaginationListVO.of(categoryVOs, categorys))
    }
}