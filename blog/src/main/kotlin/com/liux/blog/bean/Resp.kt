package com.liux.blog.bean

import org.springframework.http.HttpStatus

data class Resp<T>(
    val status: Int,
    val message: String,
    val data: T?,
) {
    companion object {
        fun succeed(): Resp<*> {
            return succeed(null)
        }

        fun <T> succeed(data: T? = null): Resp<T> {
            return Resp(HttpStatus.OK.value(), "succeed", data)
        }

        fun failed(status: Int, message: String): Resp<*> {
            return Resp(status, message, null)
        }
    }
}
