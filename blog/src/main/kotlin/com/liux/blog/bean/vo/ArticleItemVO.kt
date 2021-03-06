package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Article
import com.liux.blog.parseContent
import java.util.*

data class ArticleItemVO(
    var url: String,
    var top: Boolean,
    var title: String,
    var content: String,
    var createTime: Date,
    var updateTime: Date,
    var category: CategoryVO,
) {
    companion object {
        fun of(article: Article): ArticleItemVO {
            return ArticleItemVO(
                article.url ?: article.id.toString(),
                article.weight ?: 0 > 0,
                article.title ?: "",
                article.parseContent(),
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
                CategoryVO.of(article.category!!),
            )
        }
    }
}
