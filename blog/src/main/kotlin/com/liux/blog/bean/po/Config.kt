package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.util.*

// 配置数据
@NoArgsConstructor
data class Config(
    // 键
    var key: String,
    // 值
    var value: String,
    // 描述
    var description: String,
    // 创建时间
    var createTime: Date,
    // 更新时间
    var updateTime: Date,
)
