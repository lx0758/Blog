package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.UrlVO
import com.liux.blog.service.UrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/url")
class UrlController {

    @Autowired
    private lateinit var urlService: UrlService

    @GetMapping
    fun query(): Resp<List<UrlVO>> {
        val urls = urlService.listByAdmin()
        val urlVOs = urls.map{ UrlVO.of(it) }
        return Resp.succeed(urlVOs)
    }
}