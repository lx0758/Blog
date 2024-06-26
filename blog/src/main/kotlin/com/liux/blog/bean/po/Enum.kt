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

enum class ArticleCommentStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    DISABLE(0, "不允许"),
    ENABLE(1, "允许"),
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

enum class ArticleUrlStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    HISTORY(0, "历史的"),
    CURRENT(1, "当前的"),
    ;
}

enum class FragmentStatusEnum(
    override val value: Int,
    val description: String,
): Enum {
    DISABLE(0, "已禁用"),
    ENABLE(1, "已启用"),
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