package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Category
import java.util.*

data class CategoryVO(
    var id: Int,
    var name: String,
    var createTime: Date,
    var updateTime: Date?,
    var articleCount: Int,
) {
    companion object {
        fun of(category: Category): CategoryVO {
            return CategoryVO(
                category.id!!,
                category.name!!,
                category.createTime!!,
                category.updateTime,
                category.articleCount ?: 0,
            )
        }
    }
}
