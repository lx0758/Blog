package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.bean.vo.api.UrlVO
import com.liux.blog.service.UrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

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
}