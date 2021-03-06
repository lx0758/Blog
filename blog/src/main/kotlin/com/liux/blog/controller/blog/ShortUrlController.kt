package com.liux.blog.controller.blog

import com.liux.blog.exception.HttpNotFoundException
import com.liux.blog.service.ShortUrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.*
import javax.servlet.http.HttpServletResponse

@Controller
class ShortUrlController {

    @Autowired
    private lateinit var shortUrlService: ShortUrlService

    @GetMapping("/u/{key}")
    fun query(response: HttpServletResponse, @PathVariable("key") key: String) {
        val shortUrl = shortUrlService.getByKey(key) ?: throw HttpNotFoundException()
        response.sendRedirect(shortUrl.url)
    }
}
