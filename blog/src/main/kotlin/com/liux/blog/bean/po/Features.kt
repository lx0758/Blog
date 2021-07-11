package com.liux.blog.bean.po

import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.liux.blog.bean.NoArgsConstructor
import java.util.*

// 功能配置数据
@NoArgsConstructor
class Features(
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

    inline fun <reified T> get(): T? {
        if (value == null) return null
        return jacksonObjectMapper().readValue(value, T::class.java)
    }

    fun set(value: Any?) {
        if (value == null) {
            this.value = null
            return
        }
        this.value = jacksonObjectMapper().writeValueAsString(value)
    }

    class SMTP(
        val enable: Boolean?,
        val hostname: String?,
        val port: Int?,
        val ssl: Boolean?,
        val username: String?,
        val password: String?,
        val fromName: String?,
        val fromEmail: String?,
    ) {
        companion object {
            const val KEY = "SMTP"
        }
    }
}