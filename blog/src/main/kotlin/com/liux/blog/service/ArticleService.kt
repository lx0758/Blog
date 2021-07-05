package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Article

interface ArticleService {

    fun listByPage(pageNum: Int): Page<Article>
    fun listByArchive(pageNum: Int): Page<Article>
    fun listByCategory(categoryId: Int, pageNum: Int): Page<Article>
    fun listByTag(tagId: Int, pageNum: Int): Page<Article>
    fun listBySearch(): List<Article>
    fun listBySitemap(): List<Article>
    fun listByAdmin(
        title: String?,
        category: Int?,
        enableComment: Boolean?,
        status: Int?,
        pageNum: Int,
        pageSize: Int
    ): Page<Article>

    fun getByBlog(id: Int): Article?
    fun getByUrl(url: String): Article?
    fun getByPrev(articleId: Int): Article?
    fun getByNext(articleId: Int): Article?
    fun getByAdmin(id: Int): Article?

    fun addByAdmin(
        title: String,
        categoryId: Int,
        content: String,
        url: String?,
        weight: Int?,
        enableComment: Boolean,
        status: Int,
        tags: Array<String>?
    )

    fun updateByAdmin(
        id: Int,
        title: String,
        categoryId: Int,
        content: String,
        url: String?,
        weight: Int?,
        enableComment: Boolean,
        status: Int,
        tags: Array<String>?
    ): Int
    fun updateStatusByAdmin(id: Int, status: Int): Int

    fun deleteByAdmin(id: Int): Int
}
