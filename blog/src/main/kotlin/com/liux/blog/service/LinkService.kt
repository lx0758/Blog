package com.liux.blog.service

import com.liux.blog.bean.po.Link

interface LinkService {
    fun queryAll(): List<Link>
}
