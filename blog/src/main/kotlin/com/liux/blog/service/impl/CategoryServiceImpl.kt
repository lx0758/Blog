package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Category
import com.liux.blog.dao.CategoryMapper
import com.liux.blog.service.CategoryService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class CategoryServiceImpl: CategoryService {

    @Autowired
    private lateinit var categoryMapper: CategoryMapper

    override fun getByName(name: String): Category? {
        return categoryMapper.selectByName(name)
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
}
