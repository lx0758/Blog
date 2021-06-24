package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.User
import java.util.*

data class UserVO(
    val id: Int,
    val username: String,
    val nickname: String,
    val description: String,
    val github: String,
    val email: String,
    val lastLoginTime: Date?,
    val status: Int,
    val createTime: Date,
    val updateTime: Date,
) {
    companion object {
        fun of(user: User): UserVO {
            return UserVO(
                user.id,
                user.username!!,
                user.nickname ?: "",
                user.description ?: "",
                user.github ?: "",
                user.email ?: "",
                user.lastLoginTime,
                user.status!!,
                user.createTime!!,
                user.updateTime ?: user.createTime!!,
            )
        }
    }
}
