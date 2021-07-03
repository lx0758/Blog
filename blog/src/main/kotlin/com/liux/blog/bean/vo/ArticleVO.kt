package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Article
import com.liux.blog.parseContent
import java.util.*

data class ArticleVO(
    var id: Int,
    var url: String,
    var title: String,
    var content: String,
    var createTime: Date,
    var updateTime: Date,
    var categoryName: String,
    var authorNickname: String,
    var views: Int,
    var allowComment: Boolean,
    var tags: List<String>,
) {
    companion object {
        fun of(article: Article, catalogues: ArrayList<CatalogueVO>): ArticleVO {
            return ArticleVO(
                article.id!!,
                article.url ?: article.id.toString(),
                article.title ?: "",
                article.parseContent(catalogues),
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
                article.category!!.name!!,
                article.author!!.nickname!!,
                article.views!!,
                article.enableComment!!,
                article.tags?.map { it.name!! } ?: emptyList(),
            )
        }
    }
}
