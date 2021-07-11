package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.bean.po.Tag
import com.liux.blog.dao.TagMapper
import com.liux.blog.service.TagService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.ApplicationEventPublisher
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.util.*

@Service
class TagServiceImpl: TagService {

    @Autowired
    private lateinit var applicationEventPublisher: ApplicationEventPublisher

    @Autowired
    private lateinit var tagMapper: TagMapper

    override fun getByName(name: String): Tag? {
        return tagMapper.getByName(name)
    }

    override fun listByCount(): List<Tag> {
        return tagMapper.selectByCount()
    }

    override fun listByAdmin(name: String?, pageNum: Int, pageSize: Int): Page<Tag> {
        val page = PageHelper.startPage<Tag>(pageNum, pageSize)
        tagMapper.selectByAdmin(Tag(
            name = name,
        ))
        return page
    }

    override fun addByAdmin(name: String) {
        tagMapper.insertSelective(Tag(
            name = name,
            createTime = Date(),
            updateTime = null,
        ))
        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.TAG)
    }

    override fun updateByAdmin(id: Int, name: String): Int {
        val updateRow = tagMapper.updateByPrimaryKeySelective(Tag(
            id = id,
            name = name,
            updateTime = Date(),
        ))
        if (updateRow > 0) {
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.CATEGORY)
        }
        return updateRow
    }

    @Transactional(rollbackFor = [Exception::class])
    override fun deleteByAdmin(id: Int): Int {
        val deleteRow = tagMapper.deleteByPrimaryKey(id)
        if (deleteRow > 0) {
            tagMapper.deleteByCleanTagLink(id)
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.CATEGORY)
        }
        return deleteRow
    }
}
