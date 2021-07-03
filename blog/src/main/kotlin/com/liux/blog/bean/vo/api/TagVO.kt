package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Tag
import java.util.*

data class TagVO(
    var id: Int,
    var name: String,
    var createTime: Date,
    var updateTime: Date,
    var articleCount: Int,
) {
    companion object {
        fun of(tag: Tag): TagVO {
            return TagVO(
                tag.id!!,
                tag.name!!,
                tag.createTime!!,
                tag.updateTime ?: tag.createTime!!,
                tag.articleCount!!,
            )
        }
    }
}
