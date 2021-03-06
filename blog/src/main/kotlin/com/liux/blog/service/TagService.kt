package com.liux.blog.service

import com.liux.blog.bean.po.Tag

interface TagService {
    fun list(): List<Tag>
    fun getByName(name: String): Tag?
}
