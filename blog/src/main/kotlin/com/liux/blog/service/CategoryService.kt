package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Category

interface CategoryService {

    fun getByName(name: String): Category?

    fun listByCategory(): List<Category>
    fun listByAdmin(name: String?, pageNum: Int, pageSize: Int): Page<Category>
}
