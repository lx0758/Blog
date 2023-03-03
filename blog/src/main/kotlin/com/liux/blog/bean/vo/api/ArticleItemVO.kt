package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.ArticleCommentStatusEnum
import java.util.*

data class ArticleItemVO(
    var id: Int,
    var title: String,
    var categoryName: String,
    var url: String,
    var weight: Int,
    var views: Int,
    var enableComment: Boolean,
    var status: Int,
    var createTime: Date,
    var updateTime: Date,
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
                article.commentStatus == ArticleCommentStatusEnum.ENABLE.value,
                article.status!!,
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
            )
        }
    }
}
