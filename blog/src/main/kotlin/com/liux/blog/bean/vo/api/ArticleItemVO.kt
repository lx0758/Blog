package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.COMMENT_ENABLE
import com.liux.blog.bean.po.FORMAT_MARKDOWN
import com.liux.blog.bean.po.STATE_UNPUBLISHED
import java.util.*

data class ArticleItemVO(
    var id: Int,
    var title: String,
    var categoryName: String,
    var format: Int,
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
                article.id,
                article.title ?: "",
                article.category!!.name!!,
                article.format ?: FORMAT_MARKDOWN,
                article.url ?: "",
                article.weight ?: 0,
                article.views ?: 0,
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
                article.enableComment == COMMENT_ENABLE,
                article.status ?: STATE_UNPUBLISHED,
            )
        }
    }
}
