package com.liux.blog.service

import org.springframework.security.core.userdetails.UserDetailsService

interface UserService: UserDetailsService {
    fun refreshLastLoginTime(id: Int)
}
