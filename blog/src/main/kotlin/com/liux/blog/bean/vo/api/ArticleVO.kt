package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.ArticleCommentStatusEnum
import java.util.*

data class ArticleVO(
    var id: Int,
    var title: String,
    var categoryId: Int,
    var url: String?,
    var tags: List<String>,
    var content: String,
    var weight: Int?,
    var enableComment: Boolean,
    var status: Int,
    var createTime: Date,
    var updateTime: Date?,
) {
    companion object {
        fun of(article: Article): ArticleVO {
            return ArticleVO(
                article.id!!,
                article.title ?: "",
                article.categoryId!!,
                article.url?.url ?: "",
                article.tags?.map { it.name!! } ?: emptyList(),
                article.content ?: "",
                article.weight,
                article.commentStatus == ArticleCommentStatusEnum.ENABLE.value,
                article.status!!,
                article.createTime!!,
                article.updateTime,
            )
        }
    }
}
