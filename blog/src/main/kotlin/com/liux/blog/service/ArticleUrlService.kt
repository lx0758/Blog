package com.liux.blog.service

import com.liux.blog.bean.po.ArticleUrl

interface ArticleUrlService {
    fun findArticleUrlByPath(url: String): ArticleUrl?

    fun getCurrentArticleUrl(articleId: Int): ArticleUrl?
}