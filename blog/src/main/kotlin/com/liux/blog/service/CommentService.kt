package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Comment

interface CommentService {

    fun listByArticle(articleId: Int, pageNum: Int): Page<Comment>

    fun listByAdmin(
        articleId: Int?,
        author: String?,
        email: String?,
        ip: String?,
        content: String?,
        status: Int?,
        pageNum: Int,
        pageSize: Int
    ): Page<Comment>

    fun getCommentById(id: Int): Comment?
    fun getCountByDashboard(): Int

    fun addByBlog(
        articleId: Int,
        parentId: Int?,
        nickname: String?,
        email: String?,
        url: String?,
        content: String,
        ip: String?,
        ua: String?
    )

    fun addByAdmin(userId: Int, nickname: String, email: String, articleId: Int, parentId: Int, content: String, ip: String, ua: String)
    fun updateByAdmin(id: Int, status: Int): Int
    fun deleteByAdmin(id: Int): Int
}
