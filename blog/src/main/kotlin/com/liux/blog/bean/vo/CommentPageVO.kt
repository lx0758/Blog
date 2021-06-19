package com.liux.blog.bean.vo

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Comment

data class CommentPageVO(
    val total: Int,
    val hasMore: Boolean,
    val comments: List<CommentVO>,
) {
    companion object {
        fun of(comments: Page<Comment>): CommentPageVO {
            return CommentPageVO(
                comments.total.toInt(),
                comments.size >= comments.pageSize,
                comments.map { CommentVO.of(it) }
            )
        }
    }
}
