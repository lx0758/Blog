package com.liux.blog.service

import com.liux.blog.bean.po.ArticleUrl

interface ArticleUrlService {
    fun getUrlByPath(url: String): ArticleUrl?

    fun getCurrentUrlByArticleId(articleId: Int): ArticleUrl?
}