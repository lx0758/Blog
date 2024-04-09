package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.Fragment

interface FragmentService {
    fun apply(article: Article)
    fun listByAdmin(
        key: String?,
        content: String?,
        status: Int?,
        pageNum: Int,
        pageSize: Int,
        orderName: String?,
        orderMethod: String?
    ): Page<Fragment>
    fun getByAdmin(id: Int): Fragment?
    fun addByAdmin(userId: Int, key: String?, content: String, status: Int): Fragment
    fun updateByAdmin(id: Int, key: String?, content: String, status: Int): Int
    fun deleteByAdmin(id: Int): Int
}
