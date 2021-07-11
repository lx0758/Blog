package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.bean.po.Link
import com.liux.blog.dao.LinkMapper
import com.liux.blog.service.LinkService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.ApplicationEventPublisher
import org.springframework.stereotype.Service
import java.util.*

@Service
class LinkServiceImpl: LinkService {

    @Autowired
    private lateinit var applicationEventPublisher: ApplicationEventPublisher

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

    override fun addByAdmin(title: String, url: String, weight: Int, status: Int) {
        linkMapper.insertSelective(Link(
            title = title,
            url = url,
            weight = weight,
            status = status,
            createTime = Date(),
            updateTime = null,
        ))
        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.LINK)
    }

    override fun updateByAdmin(id: Int, title: String, url: String, weight: Int, status: Int): Int {
        val updateRow = linkMapper.updateByPrimaryKeySelective(Link(
            id = id,
            title = title,
            url = url,
            weight = weight,
            status = status,
            updateTime = Date(),
        ))
        if (updateRow > 0) {
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.LINK)
        }
        return updateRow
    }

    override fun deleteByAdmin(id: Int): Int {
        val deleteRow = linkMapper.deleteByPrimaryKey(id)
        if (deleteRow > 0) {
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.LINK)
        }
        return deleteRow
    }
}
