package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.User
import java.util.*

data class UserVO(
    val id: Int,
    val username: String,
    val avatar: String,
    val nickname: String,
    val description: String,
    val email: String,
    val accounts: User.Accounts?,
    val lastLoginTime: Date?,
    val status: Int,
    val createTime: Date,
    val updateTime: Date?,
) {
    companion object {
        fun of(user: User): UserVO {
            return UserVO(
                user.id!!,
                user.username!!,
                user.avatar ?: "",
                user.nickname ?: "",
                user.description ?: "",
                user.email ?: "",
                user.accounts,
                user.lastLoginTime,
                user.status!!,
                user.createTime!!,
                user.updateTime,
            )
        }
    }
}
