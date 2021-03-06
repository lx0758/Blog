package com.liux.blog.service

import com.liux.blog.bean.po.Category

interface CategoryService {

    fun list(): List<Category>

    fun getByName(name: String): Category?
}
