package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.UserVO
import com.liux.blog.service.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/user")
class UserController {

    @Autowired
    private lateinit var userService: UserService

    @GetMapping
    fun query(
        @RequestParam("username", required = false) username: String?,
        @RequestParam("nickname", required = false) nickname: String?,
        @RequestParam("status", required = false) status: Int?,
    ): Resp<List<UserVO>> {
        val users = userService.listByAdmin(username, nickname, status)
        val userVOs = users.map { UserVO.of(it) }
        return Resp.succeed(userVOs)
    }
}