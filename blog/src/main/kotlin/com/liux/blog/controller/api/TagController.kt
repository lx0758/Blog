package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.bean.vo.api.TagVO
import com.liux.blog.service.TagService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/tag")
class TagController {

    @Autowired
    private lateinit var tagService: TagService

    @GetMapping
    fun query(
        @RequestParam("name", required = false) name: String?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
    ): Resp<PaginationListVO<TagVO>> {
        val tags = tagService.listByAdmin(name, pageNum, pageSize)
        val tagVOs = tags.map{ TagVO.of(it) }
        return Resp.succeed(PaginationListVO.of(tagVOs, tags))
    }

    @PostMapping
    fun add(
        @RequestParam("name") name: String,
    ): Resp<*> {
        if (name.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "名称不能为空")
        }
        tagService.addByAdmin(name)
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
        val updateRow = tagService.updateByAdmin(id, name)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val deleteRow = tagService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }
}