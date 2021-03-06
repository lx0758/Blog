package com.liux.blog.bean.vo

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Comment

data class CommentVO(
    val total: Int,
    val hasMore: Boolean,
    val comments: List<CommentItemVO>,
) {
    companion object {
        fun of(comments: Page<Comment>): CommentVO {
            return CommentVO(
                comments.total.toInt(),
                comments.size >= comments.pageSize,
                comments.map { CommentItemVO.of(it) }
            )
        }
    }
}
