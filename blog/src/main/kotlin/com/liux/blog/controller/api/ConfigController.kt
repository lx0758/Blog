package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.ConfigVO
import com.liux.blog.service.ConfigService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/config")
class ConfigController {

    @Autowired
    private lateinit var configService: ConfigService

    @GetMapping
    fun query(): Resp<List<ConfigVO>> {
        val configs = configService.listByAdmin()
        val configVOs = configs.map { ConfigVO.of(it) }
        return Resp.succeed(configVOs)
    }
}