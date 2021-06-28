package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Config
import java.util.*

data class ConfigVO(
    val key: String,
    val value: String,
    val description: String,
    val createTime: Date,
    val updateTime: Date,
) {
    companion object {
        fun of(config: Config): ConfigVO {
            return ConfigVO(
                config.key!!,
                config.value ?: "",
                config.description ?: "",
                config.createTime!!,
                config.updateTime ?: config.createTime!!,
            )
        }
    }
}