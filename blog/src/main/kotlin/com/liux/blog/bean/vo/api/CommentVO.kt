package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Comment
import com.liux.blog.browser
import com.liux.blog.system
import java.util.*

data class CommentVO(
    val id: Int,
    val articleId: Int,
    val articleTitle: String,
    val nickname: String,
    val email: String,
    val url: String,
    val ip: String,
    val browser: String,
    val system: String,
    val createTime: Date,
    val updateTime: Date,
    val content: String,
    val status: Int,
) {
    companion object {
        fun of(comment: Comment): CommentVO {
            return CommentVO(
                comment.id!!,
                comment.article?.id ?: 0,
                comment.article?.title ?: "",
                comment.author ?: "匿名用户",
                comment.email ?: "",
                comment.url ?: "",
                comment.ip ?: "",
                comment.browser(),
                comment.system(),
                comment.createTime!!,
                comment.updateTime ?: comment.createTime!!,
                comment.content!!,
                comment.status!!,
            )
        }
    }
}
