package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.LinkVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.service.LinkService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

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
    ): Resp<PaginationListVO<LinkVO>> {
        val links = linkService.listByAdmin(title, url, status, pageNum, pageSize)
        val linkVOs = links.map{ LinkVO.of(it) }
        return Resp.succeed(PaginationListVO.of(linkVOs, links))
    }
}