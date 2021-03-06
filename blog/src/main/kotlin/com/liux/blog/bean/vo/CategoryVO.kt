package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Category

data class CategoryVO(
    var name: String,
    var articleCount: Int,
) {
    companion object {
        fun of(category: Category): CategoryVO {
            return CategoryVO(
                category.name!!,
                category.articleCount ?: 0,
            )
        }
    }
}
