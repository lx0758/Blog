package com.liux.blog.service.impl

import com.liux.blog.bean.po.ShortUrl
import com.liux.blog.dao.ShortUrlMapper
import com.liux.blog.service.ShortUrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class ShortUrlServiceImpl : ShortUrlService {

    @Autowired
    private lateinit var shortUrlMapper: ShortUrlMapper

    override fun getByKey(key: String): ShortUrl? {
        return shortUrlMapper.selectByKey(key)
    }
}
