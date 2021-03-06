package com.liux.blog.bean.vo

import com.liux.blog.bean.po.User

data class UserVO(
    var nickname: String,
) {
    companion object {
        fun of(user: User): UserVO {
            return UserVO(
                user.nickname ?: "无名",
            )
        }
    }
}
