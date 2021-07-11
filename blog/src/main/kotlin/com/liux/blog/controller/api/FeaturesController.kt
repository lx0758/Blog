package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.SMTPVO
import com.liux.blog.service.EmailService
import com.liux.blog.service.FeaturesService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/features")
class FeaturesController {

    @Autowired
    private lateinit var featuresService: FeaturesService
    @Autowired
    private lateinit var emailService: EmailService

    @GetMapping("smtp")
    fun querySMTP(): Resp<SMTPVO> {
        val features = featuresService.readSMTP()
        val smtpVO = SMTPVO.of(features)
        return Resp.succeed(smtpVO)
    }

    @PutMapping("smtp")
    fun updateSMTP(
        @RequestParam("enable") enable: Boolean,
        @RequestParam("hostname", required = false) hostname: String?,
        @RequestParam("port", required = false) port: Int?,
        @RequestParam("ssl", required = false) ssl: Boolean?,
        @RequestParam("username", required = false) username: String?,
        @RequestParam("password", required = false) password: String?,
        @RequestParam("fromName", required = false) fromName: String?,
        @RequestParam("fromEmail", required = false) fromEmail: String?,
    ): Resp<*> {
        val updateRow = featuresService.saveSMTP(
            enable,
            hostname,
            port,
            ssl,
            username,
            password,
            fromName,
            fromEmail,
        )
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @PostMapping("smtp/test")
    fun testSMTP(
        @RequestParam("email") email: String,
    ): Resp<*> {
        emailService.testEmail(email)
        return Resp.succeed()
    }
}