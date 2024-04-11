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
) {
    companion object {
        const val KEY_SITE_DOMAIN = "SITE_DOMAIN"
        const val KEY_SITE_TITLE = "SITE_TITLE"
        const val KEY_SITE_DESCRIPTION = "SITE_DESCRIPTION"
        const val KEY_SITE_KEYWORDS = "SITE_KEYWORDS"
        const val KEY_SITE_BEIAN_ICP = "SITE_BEIAN_ICP"
        const val KEY_SITE_BEIAN_PS = "SITE_BEIAN_PS"
        const val KEY_SITE_BAIDU = "SITE_BAIDU"
        const val KEY_SITE_CREATE_YEAR = "SITE_CREATE_YEAR"
    }
}
