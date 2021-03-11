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

    fun getArticleById(id: Int): Article?
    fun getArticleByUrl(url: String): Article?
    fun getArticleByPrev(articleId: Int): Article?
    fun getArticleByNext(articleId: Int): Article?
}
