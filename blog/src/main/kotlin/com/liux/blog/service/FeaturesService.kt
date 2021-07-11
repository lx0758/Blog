package com.liux.blog.service

import com.liux.blog.bean.po.Features

interface FeaturesService {
    fun readSMTP(): Features.SMTP?
    fun saveSMTP(
        enable: Boolean,
        hostname: String?,
        port: Int?,
        ssl: Boolean?,
        username: String?,
        password: String?,
        fromName: String?,
        fromEmail: String?
    ): Int
}