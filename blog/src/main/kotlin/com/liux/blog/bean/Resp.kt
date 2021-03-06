package com.liux.blog.bean

data class Resp<T>(
    val status: Int,
    val message: String,
    val data: T?,
) {
    companion object {
        fun succeed(): Resp<*> {
            return succeed(null)
        }

        fun <T> succeed(t: T? = null): Resp<T> {
            return Resp(200, "succeed", t)
        }

        fun failed(status: Int, message: String): Resp<*> {
            return Resp(status, message, null)
        }

        fun error(message: String): Resp<*> {
            return Resp(500, message, null)
        }
    }
}