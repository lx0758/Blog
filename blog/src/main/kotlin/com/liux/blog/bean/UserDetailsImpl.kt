package com.liux.blog.bean

import com.liux.blog.bean.po.User
import com.liux.blog.bean.po.UserStatusEnum
import org.springframework.security.core.GrantedAuthority
import org.springframework.security.core.authority.SimpleGrantedAuthority
import org.springframework.security.core.userdetails.UserDetails

class UserDetailsImpl(
    user: User,
    roles: List<String>,
) : UserDetails {

    private val id = user.id!!
    private val username = user.username
    private val password = user.password
    private val authorities = roles.map { SimpleGrantedAuthority(it) }
    private val status = user.status

    override fun getAuthorities(): List<GrantedAuthority> {
        return authorities
    }

    override fun getPassword(): String {
        return password!!
    }

    override fun getUsername(): String {
        return username!!
    }

    override fun isAccountNonExpired(): Boolean {
        return status == UserStatusEnum.ACTIVATED.value
    }

    override fun isAccountNonLocked(): Boolean {
        return true
    }

    override fun isCredentialsNonExpired(): Boolean {
        return true
    }

    override fun isEnabled(): Boolean {
        return status == UserStatusEnum.ACTIVATED.value
    }

    fun getId(): Int {
        return id
    }
}