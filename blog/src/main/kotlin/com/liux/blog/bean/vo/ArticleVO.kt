package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.STATE_ACTIVATED
import com.liux.blog.bean.po.STATE_COMMENT_ENABLE
import com.liux.blog.parseContent
import java.util.*
import kotlin.collections.ArrayList

data class ArticleVO(
    var id: Int,
    var url: String,
    var title: String,
    var content: String,
    var createTime: Date,
    var updateTime: Date,
    var category: CategoryVO,
    var author: UserVO,
    var views: Int,
    var allowComment: Boolean,
    var tags: List<String>,
    var catalogues: List<CatalogueVO>,
) {
    companion object {
        fun of(article: Article): ArticleVO {
            val catalogues = ArrayList<CatalogueVO>()
            return ArticleVO(
                article.id,
                article.url ?: article.id.toString(),
                article.title ?: "",
                article.parseContent(catalogues),
                article.createTime!!,
                article.updateTime ?: article.createTime!!,
                CategoryVO.of(article.category!!),
                UserVO.of(article.author!!),
                article.views!!,
                article.enableComment == STATE_COMMENT_ENABLE,
                article.tags?.map { it.name!! } ?: emptyList(),
                catalogues,
            )
        }
    }
}
