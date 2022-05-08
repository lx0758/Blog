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
        return urlMapper.getByKey(key)?.apply {
            views = views?.plus(1) ?: 1
            totalViews = totalViews?.plus(1) ?: 1
            urlMapper.updateByPrimaryKeySelective(Url(id, views = views, totalViews = totalViews))
        }
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
            views = 0,
            totalViews = 0,
            status = status,
            createTime = Date(),
            updateTime = null,
        ))
    }

    override fun updateByAdmin(id: Int, key: String, url: String, description: String, status: Int): Int {
        val oldUrl = urlMapper.getByPrimaryKey(id)
        var views : Int? = null
        if (oldUrl?.key != key || oldUrl.url != url) {
            views = 0
        }
        val updateRow = urlMapper.updateByPrimaryKeySelective(Url(
            id = id,
            key = key,
            url = url,
            description = description,
            views = views,
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
