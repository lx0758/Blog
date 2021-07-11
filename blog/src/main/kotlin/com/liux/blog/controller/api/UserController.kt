package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.UserVO
import com.liux.blog.service.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

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

    @GetMapping("profile")
    fun queryProfile(@CurrentUserId userId: Int): Resp<UserVO> {
        val user = userService.getById(userId)!!
        val userVO = UserVO.of(user)
        return Resp.succeed(userVO)
    }

    @PutMapping("profile")
    fun updateProfile(
        @CurrentUserId userId: Int,
        @RequestParam("avatar", required = false) avatar: String?,
        @RequestParam("nickname") nickname: String,
        @RequestParam("description", required = false) description: String?,
        @RequestParam("email") email: String,

        @RequestParam("github", required = false) github: String?,
        @RequestParam("weibo", required = false) weibo: String?,
        @RequestParam("google", required = false) google: String?,
        @RequestParam("twitter", required = false) twitter: String?,
        @RequestParam("facebook", required = false) facebook: String?,
        @RequestParam("stackOverflow", required = false) stackOverflow: String?,
        @RequestParam("youtube", required = false) youtube: String?,
        @RequestParam("instagram", required = false) instagram: String?,
        @RequestParam("skype", required = false) skype: String?,
    ): Resp<*> {
        val updateRow = userService.updateBySelf(
            userId, avatar, nickname, description, email,
            github, weibo, google, twitter, facebook, stackOverflow, youtube, instagram, skype
        )
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @PutMapping("password")
    fun updatePassword(
        @CurrentUserId userId: Int,
        @RequestParam("oldPassword") oldPassword: String,
        @RequestParam("newPassword") newPassword: String,
    ): Resp<*> {
        if (oldPassword.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "原始密码不能为空")
        }
        if (newPassword.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "新密码不能为空")
        }
        if (!newPassword.matches("^[a-zA-Z0-9.]{6,16}".toRegex())) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "新密码格式不正确")
        }
        val verify = userService.verifyPassword(userId, oldPassword)
        if (!verify) {
            throw HttpClientErrorException(HttpStatus.NOT_ACCEPTABLE, "原始密码不正确")
        }
        val update = userService.updatePassword(userId, newPassword)
        if (!update) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "修改密码失败")
        }
        return Resp.succeed()
    }
}