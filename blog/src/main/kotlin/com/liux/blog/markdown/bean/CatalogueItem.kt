package com.liux.blog.markdown.bean

data class CatalogueItem(
    var title: String,
    var anchor: String,
    var rawLevel: Int,
    var parent: CatalogueItem?,
    var childs: MutableList<CatalogueItem>? = null
)