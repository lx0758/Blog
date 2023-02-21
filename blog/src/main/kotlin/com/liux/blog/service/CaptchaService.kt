package com.liux.blog.service

import jakarta.servlet.http.HttpSession

interface CaptchaService {

    companion object {
        const val TYPE_LOGIN = 1
        const val TYPE_COMMENT = 2
    }

    fun generate(session: HttpSession, type: Int): ByteArray

    fun verify(session: HttpSession, captcha: String, type: Int, validPeriodMinute: Int = 0): Boolean
}