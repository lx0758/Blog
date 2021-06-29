package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Tag
import com.liux.blog.dao.TagMapper
import com.liux.blog.service.TagService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class TagServiceImpl: TagService {

    @Autowired
    private lateinit var tagMapper: TagMapper

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

    override fun getByName(name: String): Tag? {
        return tagMapper.selectByName(name)
    }
}
