package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.io.Serializable
import java.util.*

// 文章数据
@NoArgsConstructor
data class Fragment(
    // ID
    var id: Int? = null,
    // KEY
    var key: String? = null,
    // 内容
    var content: String? = null,
    // 作者ID
    var authorId: Int? = null,
    // 状态
    var status: Int? = null,
    // 创建时间
    var createTime: Date? = null,
    // 更新时间
    var updateTime: Date? = null,

    /* 作者 */
    var author: User? = null,
) : Serializable
