package com.liux.blog.service.impl

import com.liux.blog.bean.po.Link
import com.liux.blog.dao.LinkMapper
import com.liux.blog.service.LinkService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class LinkServiceImpl: LinkService {

    @Autowired
    private lateinit var linkMapper: LinkMapper

    override fun queryAll(): List<Link> {
        return linkMapper.select()
    }
}
