package com.liux.blog

import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.fasterxml.jackson.module.kotlin.readValue
import com.liux.blog.bean.po.Article
import com.liux.blog.bean.vo.CatalogueVO
import com.liux.blog.util.MarkdownHelper
import jakarta.servlet.http.HttpServletRequest
import org.commonmark.ext.next.FLAG_MORE_ANCHOR
import org.commonmark.ext.next.FLAG_MORE_SUSPEND
import org.commonmark.ext.next.FLAG_TOC
import org.jsoup.Jsoup
import ua_parser.Client
import java.text.SimpleDateFormat
import java.time.LocalDateTime
import java.time.ZoneId
import java.util.*


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
fun Date.toLocalDateTime(): LocalDateTime {
    return LocalDateTime.ofInstant(toInstant(), ZoneId.systemDefault())
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
        index++
    }
    if (ip.isNullOrEmpty()) {
        ip = remoteAddr
    }
    return ip!!
}
fun HttpServletRequest.getUserAgent(): String {
    return getHeader("User-Agent")
}

fun Article.render(catalogues: ArrayList<CatalogueVO>): String {
    return MarkdownHelper.parse(content, FLAG_TOC or FLAG_MORE_ANCHOR, catalogues)
}

fun Article.renderOther(): String {
    return MarkdownHelper.parse(content)
}

fun Article.renderPage(): String {
    return MarkdownHelper.parse(content, FLAG_MORE_SUSPEND)
}

fun Article.renderDescription(): String {
    val html = MarkdownHelper.parse(content, FLAG_MORE_SUSPEND)
    val document = Jsoup.parse(html)
    return document.text()
}

val Client.browser: String
    get() {
        return if (userAgent != null) "${userAgent.family} ${userAgent.major}${if (userAgent.minor.isNullOrEmpty()) "" else '.' + userAgent.minor}${if (userAgent.patch.isNullOrEmpty()) "" else '.' + userAgent.patch}" else "Unknown"
    }

val Client.system: String
    get() {
        return if (os != null) "${os.family} ${os.major}${if (os.minor.isNullOrEmpty()) "" else '.' + os.minor}${if (os.patch.isNullOrEmpty()) "" else '.' + os.patch}" else "Unknown"
    }
