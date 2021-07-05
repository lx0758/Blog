package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Article
import java.util.*

data class ArticleItemVO(
    var id: Int,
    var title: String,
    var categoryName: String,
    var url: String,
    var weight: Int,
    var views: Int,
    var createTime: Date,
    var updateTime: Date,
    var enableComment: Boolean,
    var status: Int,
) {
    companion object {
        fun of(article: Article): ArticleItemVO {
            return ArticleItemVO(
                article.id!!,
                article.title ?: "",
                article.category!!.name!!,
                article.url ?: "",
                article.weight ?: 0,
                article.views ?: 0,
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
                article.enableComment!!,
                article.status!!,
            )
        }
    }
}
