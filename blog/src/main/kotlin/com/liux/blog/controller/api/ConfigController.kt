package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.ConfigVO
import com.liux.blog.service.ConfigService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/config")
class ConfigController {

    @Autowired
    private lateinit var configService: ConfigService

    @GetMapping
    fun query(
        @RequestParam("key", required = false) key: String?,
        @RequestParam("value", required = false) value: String?,
        @RequestParam("description", required = false) description: String?,
    ): Resp<List<ConfigVO>> {
        val configs = configService.listByAdmin(key, value, description)
        val configVOs = configs.map { ConfigVO.of(it) }
        return Resp.succeed(configVOs)
    }

    @PostMapping
    fun add(
        @RequestParam("key") key: String,
        @RequestParam("value") value: String,
        @RequestParam("description") description: String,
    ): Resp<*> {
        checkParams(key, value, description)
        configService.addByAdmin(key, value, description)
        return Resp.succeed()
    }

    @PutMapping("{key}")
    fun update(
        @PathVariable("key") key: String,
        @RequestParam("value") value: String,
        @RequestParam("description") description: String,
    ): Resp<*> {
        checkParams(key, value, description)
        val updateRow = configService.updateByAdmin(key, value, description)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{key}")
    fun delete(
        @PathVariable("key") key: String,
    ): Resp<*> {
        val deleteRow = configService.deleteByAdmin(key)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }

    private fun checkParams(key: String, value: String, description: String) {
        if (key.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "键不能为空")
        }
        if (value.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "值不能为空")
        }
        if (description.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "描述不能为空")
        }
    }
}