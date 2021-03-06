package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Tag

data class TagVO(
    var name: String,
    var articleCount: Int,
) {
    companion object {
        fun of(tag: Tag): TagVO {
            return TagVO(
                tag.name!!,
                tag.articleCount!!,
            )
        }
    }
}
