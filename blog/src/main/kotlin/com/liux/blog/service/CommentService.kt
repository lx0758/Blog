package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Comment

interface CommentService {

    fun listByPage(pageNum: Int, pageSize: Int): Page<Comment>

    fun listByArticle(articleId: Int, pageNum: Int): Page<Comment>

    fun getCommentById(id: Int): Comment?

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
}