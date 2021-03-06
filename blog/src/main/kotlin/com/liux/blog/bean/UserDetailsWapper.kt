package com.liux.blog.bean

import com.liux.blog.bean.po.STATE_ACTIVATED
import com.liux.blog.bean.po.STATE_DELETED
import com.liux.blog.bean.po.User
import org.springframework.security.core.GrantedAuthority
import org.springframework.security.core.userdetails.UserDetails

class UserDetailsWapper(
    val user: User
) : UserDetails {

    override fun getAuthorities(): List<GrantedAuthority> {
        return emptyList()
    }

    override fun getPassword(): String {
        return user.password!!
    }

    override fun getUsername(): String {
        return user.username!!
    }

    override fun isAccountNonExpired(): Boolean {
        return user.status != STATE_DELETED
    }

    override fun isAccountNonLocked(): Boolean {
        return true
    }

    override fun isCredentialsNonExpired(): Boolean {
        return true
    }

    override fun isEnabled(): Boolean {
        return user.status == STATE_ACTIVATED
    }
}