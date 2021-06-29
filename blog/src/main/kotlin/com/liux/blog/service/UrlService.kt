package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Url

interface UrlService {
    fun getByKey(key: String): Url?
    fun listByAdmin(
        key: String?,
        url: String?,
        description: String?,
        status: Int?,
        pageNum: Int,
        pageSize: Int
    ): Page<Url>
}
