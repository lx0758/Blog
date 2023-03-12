package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Category

interface CategoryService {
    fun getDefaultCategory(): Category?
    fun getById(id: Int): Category?
    fun getByByName(name: String): Category?
    fun listByCategory(): List<Category>
    fun listByAdmin(
        name: String?,
        pageNum: Int,
        pageSize: Int,
        orderName: String?,
        orderMethod: String?,
    ): Page<Category>
    fun addByAdmin(name: String)
    fun updateByAdmin(id: Int, name: String): Int
    fun deleteByAdmin(id: Int): Int
}
