package com.liux.blog.service

import com.liux.blog.bean.po.ShortUrl

interface ShortUrlService {
    fun getByKey(key: String): ShortUrl?
}
