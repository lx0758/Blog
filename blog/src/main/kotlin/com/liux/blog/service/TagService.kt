package com.liux.blog.service

import com.liux.blog.bean.po.Tag

interface TagService {
    fun listByCount(): List<Tag>
    fun listByAdmin(): List<Tag>
    fun getByName(name: String): Tag?
}
