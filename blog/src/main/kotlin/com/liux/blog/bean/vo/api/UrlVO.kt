package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Url
import java.util.*

data class UrlVO(
    var id: Int,
    var key: String,
    var url: String,
    var description: String,
    var authorId: Int,
    var authorName: String,
    var createTime: Date,
    var updateTime: Date,
    var status: Int,
) {
    companion object {
        fun of(url: Url): UrlVO {
            return UrlVO(
                url.id,
                url.key!!,
                url.description!!,
                url.url!!,
                url.authorId!!,
                url.author!!.nickname!!,
                url.createTime!!,
                url.updateTime ?: url.createTime!!,
                url.status!!,
            )
        }
    }
}
