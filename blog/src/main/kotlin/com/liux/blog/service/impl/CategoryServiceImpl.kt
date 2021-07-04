package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.bean.po.Category
import com.liux.blog.dao.ArticleMapper
import com.liux.blog.dao.CategoryMapper
import com.liux.blog.service.CategoryService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.ApplicationEventPublisher
import org.springframework.stereotype.Service
import org.springframework.transaction.annotation.Transactional
import java.util.*


@Service
class CategoryServiceImpl: CategoryService {

    @Autowired
    private lateinit var applicationEventPublisher: ApplicationEventPublisher

    @Autowired
    private lateinit var categoryMapper: CategoryMapper
    @Autowired
    private lateinit var articleMapper: ArticleMapper

    override fun getDefaultCategory(): Category? {
        return categoryMapper.getDefaultCategory()
    }

    override fun getById(id: Int): Category? {
        return categoryMapper.getByPrimaryKey(id)
    }

    override fun getByByName(name: String): Category? {
        return categoryMapper.getByName(name)
    }

    override fun listByCategory(): List<Category> {
        return categoryMapper.selectByCategory()
    }

    override fun listByAdmin(name: String?, pageNum: Int, pageSize: Int): Page<Category> {
        val page = PageHelper.startPage<Category>(pageNum, pageSize)
        categoryMapper.selectByAdmin(Category(
            name = name,
        ))
        return page
    }

    override fun addByAdmin(name: String) {
        categoryMapper.insertSelective(Category(
            name = name,
            createTime = Date(),
            updateTime = null,
        ))
        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.CATEGORY)
    }

    override fun updateByAdmin(id: Int, name: String): Int {
        val updateRow = categoryMapper.updateByPrimaryKeySelective(Category(
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
        val category = getDefaultCategory()
        if (category == null || category.id == id) {
            return 0;
        }
        val deleteRow = categoryMapper.deleteByPrimaryKey(id)
        if (deleteRow > 0) {
            articleMapper.updateByMoveCategory(id, category.id!!)
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.CATEGORY)
        }
        return deleteRow
    }
}
