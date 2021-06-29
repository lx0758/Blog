package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Link

interface LinkService {
    fun listByBlog(): List<Link>
    fun listByAdmin(title: String?, url: String?, status: Int?, pageNum: Int, pageSize: Int): Page<Link>
}
