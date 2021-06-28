package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.util.*

// 友链数据
@NoArgsConstructor
data class Link(
    // ID
    var id: Int? = null,
    // 标题
    var title: String? = null,
    // 地址
    var url: String? = null,
    // 权重
    var weight: Int? = null,
    // 状态 -1_已删除 0_未启用 1_已启用
    var status: Int? = null,
    // 创建时间
    var createTime: Date? = null,
    // 更新时间
    var updateTime: Date? = null,
)
