package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.io.Serializable
import java.util.*

// 标签数据
@NoArgsConstructor
data class Tag(
    // ID
    var id: Int? = null,
    // 名称
    var name: String? = null,
    // 创建时间
    var createTime: Date? = null,

    // 文章数量
    var articleCount: Int? = null,
    // 文章列表
    var articles: List<Article>? = null,
) : Serializable
