package com.liux.blog.markdown.bean

data class CatalogueItem(
    var title: String,
    var anchor: String,
    var childs: MutableList<CatalogueItem>? = null
)