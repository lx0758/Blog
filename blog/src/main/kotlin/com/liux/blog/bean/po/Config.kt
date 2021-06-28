package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.util.*

// 配置数据
@NoArgsConstructor
data class Config(
    // 键
    var key: String? = null,
    // 值
    var value: String? = null,
    // 描述
    var description: String? = null,
    // 创建时间
    var createTime: Date? = null,
    // 更新时间
    var updateTime: Date? = null,
)
