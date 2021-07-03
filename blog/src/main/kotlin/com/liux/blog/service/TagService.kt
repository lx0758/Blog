package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Tag

interface TagService {
    fun listByCount(): List<Tag>
    fun getByName(name: String): Tag?

    fun listByAdmin(name: String?, pageNum: Int, pageSize: Int): Page<Tag>
    fun addByAdmin(name: String)
    fun updateByAdmin(id: Int, name: String): Int
    fun deleteByAdmin(id: Int): Int
}
