package com.liux.blog.bean.po

import com.liux.blog.bean.NoArgsConstructor
import java.util.*

@NoArgsConstructor
data class Upload(
    // ID
    var id: Int? = null,
    // 显示名称
    var displayName: String? = null,
    // 长度
    var length: Int? = 0,
    // 类型
    var type: String? = null,
    // 存储路径
    var path: String? = null,
    // 上传者ID
    var authorId: Int? = 0,
    // 创建时间
    var createTime: Date? = null,
    // 状态
    var status: Int? = 0,

    /* 作者 */
    var author: User? = null,
)