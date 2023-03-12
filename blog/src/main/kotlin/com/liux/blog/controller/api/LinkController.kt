package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.LinkStatusEnum
import com.liux.blog.bean.po.isValid
import com.liux.blog.bean.vo.api.LinkVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.service.LinkService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/link")
class LinkController {

    @Autowired
    private lateinit var linkService: LinkService

    @GetMapping
    fun query(
        @RequestParam("title", required = false) title: String?,
        @RequestParam("url", required = false) url: String?,
        @RequestParam("status", required = false) status: Int?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
        @RequestParam("orderName", required = false) orderName: String?,
        @RequestParam("orderMethod", required = false) orderMethod: String?,
    ): Resp<PaginationListVO<LinkVO>> {
        val links = linkService.listByAdmin(title, url, status, pageNum, pageSize, orderName, orderMethod)
        val linkVOs = links.map{ LinkVO.of(it) }
        return Resp.succeed(PaginationListVO.of(linkVOs, links))
    }

    @PostMapping
    fun add(
        @RequestParam("title") title: String,
        @RequestParam("url") url: String,
        @RequestParam("weight") weight: Int,
        @RequestParam("status") status: Int,
    ): Resp<*> {
        checkParams(title, url, weight, status)
        linkService.addByAdmin(title, url, weight, status)
        return Resp.succeed()
    }

    @PutMapping("{id}")
    fun update(
        @PathVariable("id") id: Int,
        @RequestParam("title") title: String,
        @RequestParam("url") url: String,
        @RequestParam("weight") weight: Int,
        @RequestParam("status") status: Int,
    ): Resp<*> {
        checkParams(title, url, weight, status)
        val updateRow = linkService.updateByAdmin(id, title, url, weight, status)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val deleteRow = linkService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }

    private fun checkParams(title: String, url: String, weight: Int?, status: Int) {
        if (title.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "键不能为空")
        }
        if (url.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "键不能为空")
        }
        if (!LinkStatusEnum.values().isValid(status)) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "状态不正确")
        }
    }
}
