package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Article
import java.util.*

data class ArticleVO(
    var id: Int,
    var title: String,
    var categoryId: Int,
    var url: String?,
    var tags: List<String>,
    var content: String,
    var weight: Int,
    var enableComment: Boolean,
    var status: Int,
    var createTime: Date,
    var updateTime: Date,
) {
    companion object {
        fun of(article: Article): ArticleVO {
            return ArticleVO(
                article.id!!,
                article.title ?: "",
                article.categoryId!!,
                article.url,
                article.tags?.map { it.name!! } ?: emptyList(),
                article.content ?: "",
                article.weight ?: 0,
                article.enableComment ?: false,
                article.status!!,
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
            )
        }
    }
}
