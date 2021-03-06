package com.liux.blog.service.impl

import com.liux.blog.bean.UserDetailsWapper
import com.liux.blog.bean.po.User
import com.liux.blog.dao.UserMapper
import com.liux.blog.service.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.security.core.userdetails.UserDetails
import org.springframework.security.core.userdetails.UsernameNotFoundException
import org.springframework.stereotype.Service
import java.util.*

@Service
class UserServiceImpl: UserService {

    @Autowired
    private lateinit var userMapper: UserMapper

    override fun refreshLastLoginTime(id: Int) {
        userMapper.updateByPrimaryKeySelective(User(id, lastLoginTime = Date()))
    }

    @Throws(UsernameNotFoundException::class)
    override fun loadUserByUsername(username: String): UserDetails {
        val user = userMapper.selectByUsername(username) ?: throw UsernameNotFoundException("Account does not exist")
        return UserDetailsWapper(user)
    }
}
