package com.liux.blog.service

import com.liux.blog.bean.po.User
import org.springframework.security.core.userdetails.UserDetailsService

interface UserService: UserDetailsService {
    fun getById(id: Int): User?
    fun listByAdmin(username: String?, nickname: String?, status: Int?): List<User>
    fun refreshLastLoginTime(id: Int)
}
