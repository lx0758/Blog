package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Article
import com.liux.blog.renderPage
import java.util.*

data class ArticlePageVO(
    var url: String,
    var top: Boolean,
    var title: String,
    var content: String,
    var createTime: Date,
    var updateTime: Date,
    var categoryName: String,
) {
    companion object {
        fun of(article: Article): ArticlePageVO {
            return ArticlePageVO(
                article.url ?: article.id.toString(),
                (article.weight ?: 0) > 0,
                article.title ?: "",
                article.renderPage(),
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
                article.category!!.name!!,
            )
        }
    }
}
