package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.CategoryVO
import com.liux.blog.service.CategoryService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/category")
class CategoryController {

    @Autowired
    private lateinit var categoryService: CategoryService

    @GetMapping
    fun query(): Resp<List<CategoryVO>> {
        val categorys = categoryService.list()
        val categoryVOs = categorys.map { CategoryVO.of(it) }
        return Resp.succeed(categoryVOs)
    }
}