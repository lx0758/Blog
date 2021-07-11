package com.liux.blog.service

interface EmailService {

    fun testEmail(
        to: String,
    )

    fun sendCommentAddEmail(
        articleId: Int,
        nickname: String,
        content: String,
    )

    fun sendCommentReplayEmail(
        articleId: Int,
        commentId: Int,
        content: String,
    )
}