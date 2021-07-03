package com.liux.blog.bean.po

interface Enum {
    val value: Int
}

inline fun <reified T: Enum> Array<T>.enum(statusValue: Int): T? {
    for (status in this) {
        if (status.value == statusValue) return status
    }
    return null
}

inline fun <reified T: Enum> Array<T>.isValid(statusValue: Int): Boolean {
    return enum(statusValue) != null
}

enum class UserStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    NOT_ACTIVATED(0, "未激活"),
    ACTIVATED(1, "已激活"),
    ;
}

enum class ArticleFormatEnum(
    override val value: Int,
    val description: String,
): Enum {
    MARKDOWN(0, "Markdown"),
    HTML(1, "Html"),
    ;
}

enum class ArticleStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    UNPUBLISHED(0, "草稿"),
    PUBLISHED(1, "已发布"),
    ;
}

enum class CommentStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    PENDING(0, "待审核"),
    AUDITED(1, "已审核"),
    ;
}

enum class LinkStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    DISABLE(0, "已禁用"),
    ENABLE(1, "已启用"),
    ;
}

enum class UrlStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    DISABLE(0, "已禁用"),
    ENABLE(1, "已启用"),
    ;
}