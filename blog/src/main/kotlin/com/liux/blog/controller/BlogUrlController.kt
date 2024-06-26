package com.liux.blog.controller

import com.liux.blog.service.UrlService
import jakarta.servlet.http.HttpServletResponse
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.client.HttpClientErrorException

@Controller
class BlogUrlController {

    @Autowired
    private lateinit var urlService: UrlService

    @GetMapping("/u/{key}")
    fun query(response: HttpServletResponse, @PathVariable("key") key: String) {
        val url = urlService.getByKey(key) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "短链不存在")
        response.sendRedirect(url.url)
    }
}
