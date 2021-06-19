package com.liux.blog.controller

import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.GetMapping

@Controller
class AdminController {

    @GetMapping("/admin")
    fun index(): String {
        return "redirect:/admin/index.html"
    }
}