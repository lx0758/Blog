package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.ArticleCommentStatusEnum
import com.liux.blog.render
import com.liux.blog.renderDescription
import java.util.*

data class ArticleVO(
    var id: Int,
    var url: String,
    var title: String,
    var description: String,
    var content: String,
    var createTime: Date,
    var updateTime: Date,
    var categoryName: String,
    var authorNickname: String,
    var views: Int,
    var enableComment: Boolean,
    var tags: List<String>,
    var catalogues: List<CatalogueVO>,
) {
    companion object {
        fun of(article: Article): ArticleVO {
            val catalogues = ArrayList<CatalogueVO>()
            return ArticleVO(
                article.id!!,
                article.url?.url ?: article.id.toString(),
                article.title ?: "",
                article.renderDescription(),
                article.render(catalogues),
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
                article.category!!.name!!,
                article.author!!.nickname!!,
                article.views!!,
                article.commentStatus == ArticleCommentStatusEnum.ENABLE.value,
                article.tags?.map { it.name!! } ?: emptyList(),
                catalogues,
            )
        }
    }
}
