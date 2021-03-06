package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Comment
import com.liux.blog.toDateString
import ua_parser.Parser

data class CommentItemVO(
    val id: Int,
    val avatar: String,
    val nickname: String,
    val browser: String,
    val system: String,
    val time: String,
    val content: String,
    val childs: List<CommentItemVO>?,
) {
    companion object {
        val PARSER = Parser()
        fun of(comment: Comment): CommentItemVO {
            val client = PARSER.parse(comment.ua)
            return CommentItemVO(
                comment.id,
                "/blog/images/avatar.gif",
                comment.author ?: "匿名用户",
                if (client.userAgent != null) "${client.userAgent.family} ${client.userAgent.major}${if (client.userAgent.minor.isNullOrEmpty()) "" else '.' + client.userAgent.minor}${if (client.userAgent.patch.isNullOrEmpty()) "" else '.' + client.userAgent.patch}" else "Unknown",
                if (client.os != null) "${client.os.family} ${client.os.major}${if (client.os.minor.isNullOrEmpty()) "" else '.' + client.os.minor}${if (client.os.patch.isNullOrEmpty()) "" else '.' + client.os.patch}" else "Unknown",
                comment.createTime!!.toDateString(),
                comment.content!!,
                comment.childs?.map { of(it) }
            )
        }
    }
}
