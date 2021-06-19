package com.liux.blog.service.impl

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

    override fun listByAdmin(): List<Tag> {
        return tagMapper.selectByCount()
    }

    override fun getByName(name: String): Tag? {
        return tagMapper.selectByName(name)
    }
}
