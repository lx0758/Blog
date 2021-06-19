package com.liux.blog.controller

import com.liux.blog.exception.HttpNotFoundException
import com.liux.blog.service.UrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.*
import javax.servlet.http.HttpServletResponse

@Controller
class BlogUrlController {

    @Autowired
    private lateinit var urlService: UrlService

    @GetMapping("/u/{key}")
    fun query(response: HttpServletResponse, @PathVariable("key") key: String) {
        val url = urlService.getByKey(key) ?: throw HttpNotFoundException()
        response.sendRedirect(url.url)
    }
}
