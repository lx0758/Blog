package com.liux.blog.controller.api

import com.liux.blog.service.CaptchaService
import jakarta.servlet.http.HttpServletRequest
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.MediaType
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/captcha")
class CaptchaController {

    @Autowired
    private lateinit var captchaService: CaptchaService

    @GetMapping(produces = [MediaType.IMAGE_JPEG_VALUE])
    fun captcha(request: HttpServletRequest): ByteArray {
        return captchaService.generate(request.session, CaptchaService.TYPE_LOGIN)
    }
}
