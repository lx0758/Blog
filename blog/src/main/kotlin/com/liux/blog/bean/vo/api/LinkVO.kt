package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Link
import java.util.*

data class LinkVO(
    var id: Int,
    var title: String,
    var url: String,
    var weight: Int,
    var createTime: Date,
    var updateTime: Date?,
    var status: Int,
) {
    companion object {
        fun of(link: Link): LinkVO {
            return LinkVO(
                link.id!!,
                link.title!!,
                link.url!!,
                link.weight!!,
                link.createTime!!,
                link.updateTime,
                link.status!!,
            )
        }
    }
}
