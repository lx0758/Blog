package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.util.*

// 短连接数据
@NoArgsConstructor
class Url(
    // ID
    var id: Int? = null,
    // 键
    var key: String? = null,
    // 链接地址
    var url: String? = null,
    // 描述信息
    var description: String? = null,
    // 作者ID
    var authorId: Int? = null,
    // 创建时间
    var createTime: Date? = null,
    // 更新时间
    var updateTime: Date? = null,
    // 状态 -1_已删除 0_禁用 1_启用
    var status: Int? = null,

    // 作者
    var author: User? = null
)
