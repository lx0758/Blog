package com.liux.blog.controller.admin.api

import com.liux.blog.bean.Resp
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/admin/api/overview")
class OverviewController {

    @GetMapping
    fun get(): Resp<*> {
        return Resp.succeed("Hello Word!")
    }
}