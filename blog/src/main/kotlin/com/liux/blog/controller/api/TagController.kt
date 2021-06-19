package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.TagVO
import com.liux.blog.service.TagService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/tag")
class TagController {

    @Autowired
    private lateinit var tagService: TagService

    @GetMapping
    fun query(): Resp<List<TagVO>> {
        val tags = tagService.listByAdmin()
        val tagVOs = tags.map{ TagVO.of(it) }
        return Resp.succeed(tagVOs)
    }
}