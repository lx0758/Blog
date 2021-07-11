package com.liux.blog.service

import com.liux.blog.bean.po.User
import org.springframework.security.core.userdetails.UserDetailsService

interface UserService: UserDetailsService {
    fun getByOwner(): User
    fun getById(id: Int): User?
    fun listByAdmin(username: String?, nickname: String?, status: Int?): List<User>
    fun updateByLogin(id: Int)
    fun updateBySelf(
        id: Int,
        avatar: String?,
        nickname: String,
        description: String?,
        email: String,

        github: String?,
        weibo: String?,
        google: String?,
        twitter: String?,
        facebook: String?,
        stackOverflow: String?,
        youtube: String?,
        instagram: String?,
        skype: String?,
    ): Int

    fun verifyPassword(id: Int, password: String): Boolean
    fun updatePassword(id: Int, password: String): Boolean
}
