package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.util.*

// 文章URL数据
@NoArgsConstructor
data class ArticleUrl(
    // URL
    var url: String? = null,
    // 文章ID
    var articleId: Int? = null,
    // URL状态 0_历史的 1_当前的
    var status: Int? = null,
    // 创建时间
    var createTime: Date? = null,
)