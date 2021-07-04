package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.io.Serializable
import java.util.*

// 用户数据
@NoArgsConstructor
data class User(
    // ID
    var id: Int? = null,
    // 用户名
    var username: String? = null,
    // 用户密码
    var password: String? = null,
    // 头像
    var avatar: String? = null,
    // 昵称
    var nickname: String? = null,
    // 个性签名
    var description: String? = null,
    // 电子邮箱
    var email: String? = null,
    // 社交账号
    var accounts: Accounts? = null,
    // 最后登录时间
    var lastLoginTime: Date? = null,
    // 状态
    var status: Int? = null,
    // 创建时间
    var createTime: Date? = null,
    // 更新时间
    var updateTime: Date? = null,
) : Serializable {
    data class Accounts(
        val github: String?,
        val weibo: String?,
        val google: String?,
        val twitter: String?,
        val facebook: String?,
        val stackOverflow: String?,
        val youtube: String?,
        val instagram: String?,
        val skype: String?,
    ) : Serializable
}
