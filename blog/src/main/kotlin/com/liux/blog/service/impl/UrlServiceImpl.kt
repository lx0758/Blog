package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Url
import com.liux.blog.dao.UrlMapper
import com.liux.blog.service.UrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import java.util.*

@Service
class UrlServiceImpl : UrlService {

    @Autowired
    private lateinit var urlMapper: UrlMapper

    override fun getByKey(key: String): Url? {
        return urlMapper.getByKey(key)
    }

    override fun listByAdmin(
        key: String?,
        url: String?,
        description: String?,
        status: Int?,
        pageNum: Int,
        pageSize: Int
    ): Page<Url> {
        val page = PageHelper.startPage<Url>(pageNum, pageSize)
        urlMapper.selectByAdmin(Url(
            key = key,
            url = url,
            description = description,
            status = status,
        ))
        return page
    }

    override fun addByAdmin(userId: Int, key: String, url: String, description: String, status: Int) {
        urlMapper.insertSelective(Url(
            key = key,
            url = url,
            description = description,
            authorId = userId,
            status = status,
            createTime = Date(),
            updateTime = null,
        ))
    }

    override fun updateByAdmin(id: Int, key: String, url: String, description: String, status: Int): Int {
        val updateRow = urlMapper.updateByPrimaryKeySelective(Url(
            id = id,
            key = key,
            url = url,
            description = description,
            status = status,
            updateTime = Date(),
        ))
        return updateRow
    }

    override fun deleteByAdmin(id: Int): Int {
        val deleteRow = urlMapper.deleteByPrimaryKey(id)
        return deleteRow
    }
}
