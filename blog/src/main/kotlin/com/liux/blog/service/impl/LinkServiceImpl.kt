package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Link
import com.liux.blog.dao.LinkMapper
import com.liux.blog.service.LinkService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class LinkServiceImpl: LinkService {

    @Autowired
    private lateinit var linkMapper: LinkMapper

    override fun listByBlog(): List<Link> {
        return linkMapper.selectByBlog()
    }

    override fun listByAdmin(title: String?, url: String?, status: Int?, pageNum: Int, pageSize: Int): Page<Link> {
        val page = PageHelper.startPage<Link>(pageNum, pageSize)
        linkMapper.selectByAdmin(Link(
            title = title,
            url = url,
            status = status,
        ))
        return page
    }
}
