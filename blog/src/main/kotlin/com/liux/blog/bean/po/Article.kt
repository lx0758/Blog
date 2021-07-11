package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.io.Serializable
import java.util.*

// 文章数据
@NoArgsConstructor
data class Article(
    // ID
    var id: Int? = null,
    // 标题
    var title: String? = null,
    // 内容
    var content: String? = null,
    // 分类ID
    var categoryId: Int? = null,
    // 作者ID
    var authorId: Int? = null,
    // 自定义url
    var url: String? = null,
    // 权重
    var weight: Int? = null,
    // 阅读量
    var views: Int? = null,
    // 允许评论
    var enableComment: Boolean? = null,
    // 状态
    var status: Int? = null,
    // 创建时间
    var createTime: Date? = null,
    // 更新时间
    var updateTime: Date? = null,

    /* 分类 */
    var category: Category? = null,
    /* 作者 */
    var author: User? = null,
    /* 标签 */
    var tags: List<Tag>? = null,
) : Serializable
