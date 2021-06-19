package com.liux.blog.service

import com.liux.blog.bean.po.Url

interface UrlService {
    fun getByKey(key: String): Url?
    fun listByAdmin(): List<Url>
}
