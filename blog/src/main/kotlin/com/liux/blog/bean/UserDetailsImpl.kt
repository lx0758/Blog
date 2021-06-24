package com.liux.blog.bean

import com.liux.blog.bean.po.STATE_ACTIVATED
import com.liux.blog.bean.po.STATE_DELETED
import com.liux.blog.bean.po.User
import org.springframework.security.core.GrantedAuthority
import org.springframework.security.core.userdetails.UserDetails

class UserDetailsImpl(
    user: User
) : UserDetails {

    private val id = user.id
    private val username = user.username
    private val password = user.password
    private val status = user.status

    override fun getAuthorities(): List<GrantedAuthority> {
        return emptyList()
    }

    override fun getPassword(): String {
        return password!!
    }

    override fun getUsername(): String {
        return username!!
    }

    override fun isAccountNonExpired(): Boolean {
        return status != STATE_DELETED
    }

    override fun isAccountNonLocked(): Boolean {
        return true
    }

    override fun isCredentialsNonExpired(): Boolean {
        return true
    }

    override fun isEnabled(): Boolean {
        return status == STATE_ACTIVATED
    }

    fun getId(): Int {
        return id
    }
}