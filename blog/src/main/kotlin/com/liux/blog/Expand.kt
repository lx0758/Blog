package com.liux.blog

import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import com.fasterxml.jackson.module.kotlin.readValue
import com.liux.blog.bean.po.Article
import com.liux.blog.bean.po.ArticleFormatEnum
import com.liux.blog.bean.po.Comment
import com.liux.blog.bean.vo.CatalogueVO
import com.liux.blog.config.KaptchaConfig
import com.liux.blog.markdown.MarkdownHelper
import ua_parser.Parser
import java.text.SimpleDateFormat
import java.time.LocalDateTime
import java.time.ZoneId
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
fun HttpServletRequest.checkVerifyCode(verifyCode: String, validPeriodMinute: Int = 0): Boolean {
    val sessionVerifyCode = session.getAttribute(KaptchaConfig.KAPTCHA_SESSION_CONFIG_KEY) ?: return false
    if (sessionVerifyCode != verifyCode) return false
    if (validPeriodMinute > 0) {
        val sessionVerifyCodeDate = session.getAttribute(KaptchaConfig.KAPTCHA_SESSION_CONFIG_DATE) ?: return false
        if (sessionVerifyCodeDate !is Date) return false
        if (!sessionVerifyCodeDate.toLocalDateTime().plusMinutes(validPeriodMinute.toLong()).isAfter(LocalDateTime.now())) return false
    }
    session.removeAttribute(KaptchaConfig.KAPTCHA_SESSION_CONFIG_KEY)
    session.removeAttribute(KaptchaConfig.KAPTCHA_SESSION_CONFIG_DATE)
    return true
}

fun Article.parseContent(catalogues: ArrayList<CatalogueVO>? = null): String {
    val localContent = content
    if (localContent.isNullOrEmpty()) return ""
    return when (format) {
        ArticleFormatEnum.MARKDOWN.value -> {
            MarkdownHelper.parse(localContent, catalogues)
        }
        ArticleFormatEnum.HTML.value -> {
            localContent
        }
        else -> ""
    }
}

val PARSER = Parser()
fun Comment.browser(): String {
    val client = PARSER.parse(ua)
    return if (client.userAgent != null) "${client.userAgent.family} ${client.userAgent.major}${if (client.userAgent.minor.isNullOrEmpty()) "" else '.' + client.userAgent.minor}${if (client.userAgent.patch.isNullOrEmpty()) "" else '.' + client.userAgent.patch}" else "Unknown"
}
fun Comment.system(): String {
    val client = PARSER.parse(ua)
    return if (client.os != null) "${client.os.family} ${client.os.major}${if (client.os.minor.isNullOrEmpty()) "" else '.' + client.os.minor}${if (client.os.patch.isNullOrEmpty()) "" else '.' + client.os.patch}" else "Unknown"
}
