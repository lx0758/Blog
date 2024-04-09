package com.liux.blog.bean.vo

data class CatalogueVO(
    var title: String,
    var anchor: String,
    var childs: MutableList<CatalogueVO>? = null
)
