package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Comment
import com.liux.blog.browser
import com.liux.blog.system
import com.liux.blog.toDateString

data class CommentVO(
    val id: Int,
    val avatar: String,
    val nickname: String,
    val browser: String,
    val system: String,
    val time: String,
    val content: String,
    val children: List<CommentVO>?,
) {
    companion object {
        fun of(comment: Comment): CommentVO {
            return CommentVO(
                comment.id!!,
                "/blog/images/avatar.gif",
                comment.author ?: "匿名用户",
                comment.browser(),
                comment.system(),
                comment.createTime!!.toDateString(),
                comment.content!!,
                comment.childs?.map { of(it) }
            )
        }
    }
}
