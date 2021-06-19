package com.liux.blog.controller.admin

import com.liux.blog.bean.Resp
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/url")
class UrlController {

    @GetMapping
    fun query(): Resp<*> {
        return Resp.succeed("Hello Word!")
    }
}