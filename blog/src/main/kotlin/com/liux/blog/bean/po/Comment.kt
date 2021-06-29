package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.io.Serializable
import java.util.*

// 评论数据
@NoArgsConstructor
data class Comment(
    // ID
    var id: Int? = null,
    // 文章ID
    var articleId: Int? = null,
    // 父评论ID
    var parentId: Int? = null,
    // 作者
    var author: String? = null,
    // 作者
    var authorId: String? = null,
    // 电子邮箱
    var email: String? = null,
    // 链接地址
    var url: String? = null,
    // 内容
    var content: String? = null,
    // IP地址
    var ip: String? = null,
    // 浏览器标识
    var ua: String? = null,
    // 状态 -1_已删除 0_未审核 1_已审核
    var status: Int? = null,
    // 创建时间
    var createTime: Date? = null,
    // 更新时间
    var updateTime: Date? = null,

    // 文章
    var article: Article? = null,
    // 子评论
    var childs: List<Comment>? = null
) : Serializable
