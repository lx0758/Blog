package com.liux.blog

import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.fasterxml.jackson.module.kotlin.readValue
import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.FORMAT_HTML
import com.liux.blog.bean.po.FORMAT_MARKDOWN
import com.liux.blog.bean.vo.CatalogueVO
import com.liux.blog.markdown.MarkdownHelper
import java.text.SimpleDateFormat
import java.util.*

val JSON = jacksonObjectMapper()

fun Any.toJSONString(): String {
    return JSON.writeValueAsString(this)
}

inline fun <reified T> String.toBean(): T {
    return JSON.readValue(this)
}

fun Date.toDateString(pattern: String = "yyyy-MM-dd HH:mm:ss"): String {
    val simpleDateFormat = SimpleDateFormat(pattern)
    return simpleDateFormat.format(this)
}

fun Article.parseContent(catalogues: MutableList<CatalogueVO>? = null): String {
    val localContent = content
    if (localContent.isNullOrBlank()) return ""
    return when (format) {
        FORMAT_MARKDOWN -> {
            MarkdownHelper.parse(localContent, catalogues)
        }
        FORMAT_HTML -> {
            localContent
        }
        else -> ""
    }
}
