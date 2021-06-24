package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.User
import com.liux.blog.bean.vo.api.UserVO
import com.liux.blog.service.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/profile")
class ProfileController {

    @Autowired
    private lateinit var userService: UserService

    @GetMapping
    fun query(@CurrentUserId userId: Int): Resp<UserVO> {
        val user = userService.getById(userId)!!
        val userVO = UserVO.of(user)
        return Resp.succeed(userVO)
    }
}