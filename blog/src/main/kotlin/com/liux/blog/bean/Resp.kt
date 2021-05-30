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

        fun succeed(message: String): Resp<*> {
            return succeed(message, null)
        }

        fun <T> succeed(data: T? = null): Resp<T> {
            return succeed("succeed", data)
        }

        fun <T> succeed(message: String, data: T?): Resp<T> {
            return Resp(200, message, data)
        }

        fun failed(message: String): Resp<*> {
            return failed(500, message)
        }

        fun failed(status: Int, message: String): Resp<*> {
            return failed(status, message, null)
        }

        fun <T> failed(status: Int, message: String, data: T?): Resp<T> {
            return Resp(status, message, data)
        }
    }
}
