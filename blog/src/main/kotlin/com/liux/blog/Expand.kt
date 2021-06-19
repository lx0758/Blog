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
import javax.servlet.http.HttpServletRequest

fun Any.toJSONString(): String {
    return jacksonObjectMapper().writeValueAsString(this)
}

fun Any.toFormatJSONString(): String {
    return jacksonObjectMapper().writerWithDefaultPrettyPrinter().writeValueAsString(this)
}

inline fun <reified T> String.toBean(): T {
    return jacksonObjectMapper().readValue(this)
}

fun Date.toDateString(pattern: String = "yyyy-MM-dd HH:mm:ss"): String {
    val simpleDateFormat = SimpleDateFormat(pattern)
    return simpleDateFormat.format(this)
}

val IP_HEADERS = arrayOf(
    "X-Forwarded-For",
    "Proxy-Client-IP",
    "WL-Proxy-Client-IP",
    "HTTP_CLIENT_IP",
    "X-Real-IP",
)
fun HttpServletRequest.getIp(): String {
    var index = 0
    var ip: String? = null
    while (ip.isNullOrEmpty() && index < IP_HEADERS.size) {
        ip = getHeader(IP_HEADERS[index])
        index ++
    }
    if (ip.isNullOrEmpty()) {
        ip = remoteAddr
    }
    return ip!!
}

fun HttpServletRequest.getUserAgent(): String {
    return getHeader("User-Agent")
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
