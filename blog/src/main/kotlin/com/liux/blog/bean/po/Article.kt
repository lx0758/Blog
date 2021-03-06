package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.io.Serializable
import java.util.*

// 文章数据
@NoArgsConstructor
data class Article(
    // ID
    var id: Int,
    // 标题
    var title: String? = null,
    // 内容
    var content: String? = null,
    // 内容格式 0_markdown 1_html
    var format: Int? = null,
    // 自定义url
    var url: String? = null,
    // 权重
    var weight: Int? = null,
    // 阅读量
    var views: Int? = null,
    // 允许评论 0_禁止评论 1_允许评论
    var enableComment: Int? = null,
    // 状态 -1_已删除 0_草稿 1_发布
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
