package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.CategoryVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.service.CategoryService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

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
        @RequestParam("orderName", required = false) orderName: String?,
        @RequestParam("orderMethod", required = false) orderMethod: String?,
    ): Resp<PaginationListVO<CategoryVO>> {
        val categorys = categoryService.listByAdmin(name, pageNum, pageSize, orderName, orderMethod)
        val categoryVOs = categorys.map { CategoryVO.of(it) }
        return Resp.succeed(PaginationListVO.of(categoryVOs, categorys))
    }

    @PostMapping
    fun add(
        @RequestParam("name") name: String,
    ): Resp<*> {
        if (name.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "名称不能为空")
        }
        categoryService.addByAdmin(name)
        return Resp.succeed()
    }

    @PutMapping("{id}")
    fun update(
        @PathVariable("id") id: Int,
        @RequestParam("name") name: String,
    ): Resp<*> {
        if (name.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "名称不能为空")
        }
        val updateRow = categoryService.updateByAdmin(id, name)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val category = categoryService.getDefaultCategory() ?: throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "默认分类不存在")
        if (category.id == id) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "默认分类不能删除")
        }
        val deleteRow = categoryService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }
}