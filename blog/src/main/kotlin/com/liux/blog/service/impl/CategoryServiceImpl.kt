package com.liux.blog.service.impl

import com.liux.blog.bean.po.Category
import com.liux.blog.dao.CategoryMapper
import com.liux.blog.service.CategoryService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class CategoryServiceImpl: CategoryService {

    @Autowired
    private lateinit var categoryMapper: CategoryMapper

    override fun list(): List<Category> {
        return categoryMapper.select()
    }

    override fun getByName(name: String): Category? {
        return categoryMapper.selectByName(name)
    }
}
