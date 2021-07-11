package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.UrlStatusEnum
import com.liux.blog.bean.po.isValid
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.bean.vo.api.UrlVO
import com.liux.blog.service.UrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/url")
class UrlController {

    @Autowired
    private lateinit var urlService: UrlService

    @GetMapping
    fun query(
        @RequestParam("key", required = false) key: String?,
        @RequestParam("url", required = false) url: String?,
        @RequestParam("description", required = false) description: String?,
        @RequestParam("status", required = false) status: Int?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
    ): Resp<PaginationListVO<UrlVO>> {
        val urls = urlService.listByAdmin(key, url, description, status, pageNum, pageSize)
        val urlVOs = urls.map{ UrlVO.of(it) }
        return Resp.succeed(PaginationListVO.of(urlVOs, urls))
    }

    @PostMapping
    fun add(
        @CurrentUserId userId: Int,
        @RequestParam("key") key: String,
        @RequestParam("url") url: String,
        @RequestParam("description") description: String,
        @RequestParam("status") status: Int,
    ): Resp<*> {
        checkParams(key, url, description, status)
        urlService.addByAdmin(userId, key, url, description, status)
        return Resp.succeed()
    }

    @PutMapping("{id}")
    fun update(
        @PathVariable("id") id: Int,
        @RequestParam("key") key: String,
        @RequestParam("url") url: String,
        @RequestParam("description") description: String,
        @RequestParam("status") status: Int,
    ): Resp<*> {
        checkParams(key, url, description, status)
        val updateRow = urlService.updateByAdmin(id, key, url, description, status)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val deleteRow = urlService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }

    private fun checkParams(key: String, url: String, description: String, status: Int) {
        if (key.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "名称不能为空")
        }
        if (url.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "链接不能为空")
        }
        if (description.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "描述不能为空")
        }
        if (!UrlStatusEnum.values().isValid(status)) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "状态不正确")
        }
    }
}