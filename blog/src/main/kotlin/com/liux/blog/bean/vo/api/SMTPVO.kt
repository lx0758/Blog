package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Features

class SMTPVO(
    val enable: Boolean,
    val hostname: String?,
    val port: Int?,
    val ssl: Boolean?,
    val username: String?,
    val password: String?,
    val fromName: String?,
    val fromEmail: String?,
) {
    companion object {
        fun of(smtp: Features.SMTP?):SMTPVO {
            return SMTPVO(
                smtp?.enable ?: false,
                smtp?.hostname,
                smtp?.port,
                smtp?.ssl,
                smtp?.username,
                smtp?.password,
                smtp?.fromName,
                smtp?.fromEmail,
            )
        }
    }
}