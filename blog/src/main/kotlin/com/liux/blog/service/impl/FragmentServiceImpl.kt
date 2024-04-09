package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.Fragment
import com.liux.blog.bean.po.Tag
import com.liux.blog.dao.FragmentMapper
import com.liux.blog.service.FragmentService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import java.util.*

@Service
class FragmentServiceImpl : FragmentService {

    @Autowired
    private lateinit var fragmentMapper: FragmentMapper

    companion object {
        private val regex = "<fragment include=['\"]([A-Za-z0-9-]+)['\"]/>".toRegex()
    }

    override fun apply(article: Article) {
        var content = article.content ?: return
        val sequence = regex.findAll(content)
        sequence.forEach {
            val includeTag = it.value
            val fragmentValue = it.groupValues[1]
            val fragmentId = try {
                fragmentValue.toInt()
            } catch (e: NumberFormatException) {
                0
            }
            val fragmentContent = fragmentMapper.getByIdOrKey(fragmentId, fragmentValue)?.content ?: ""
            content = content.replace(includeTag, fragmentContent)
        }
        article.content = content
    }

    override fun listByAdmin(
        key: String?,
        content: String?,
        status: Int?,
        pageNum: Int,
        pageSize: Int,
        orderName: String?,
        orderMethod: String?
    ): Page<Fragment> {
        val page = PageHelper.startPage<Fragment>(pageNum, pageSize)
        fragmentMapper.selectByAdmin(Fragment(
            key = key,
            content = content,
            status = status,
        ))
        return page
    }

    override fun getByAdmin(id: Int): Fragment? {
        return fragmentMapper.getByAdmin(id)
    }

    override fun addByAdmin(userId: Int, key: String?, content: String, status: Int): Fragment {
        val fragment = Fragment(
            key = key,
            content = content,
            authorId = userId,
            status = status,
            createTime = Date(),
            updateTime = null,
        )
        fragmentMapper.insertSelective(fragment)
        return fragment
    }

    override fun updateByAdmin(id: Int, key: String?, content: String, status: Int): Int {
        return fragmentMapper.updateByPrimaryKeySelective(Fragment(
            id = id,
            key = key,
            content = content,
            status = status,
            updateTime = Date(),
        ))
    }

    override fun deleteByAdmin(id: Int): Int {
        return fragmentMapper.deleteByPrimaryKey(id)
    }
}
