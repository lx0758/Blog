package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Article
import java.util.*

data class ArticleArchiveVO(
    var url: String,
    var title: String,
    var createTime: Date,
) {
    companion object {
        fun of(article: Article): ArticleArchiveVO {
            return ArticleArchiveVO(
                article.url ?: article.id.toString(),
                article.title ?: "",
                article.createTime!!,
            )
        }
    }
}
