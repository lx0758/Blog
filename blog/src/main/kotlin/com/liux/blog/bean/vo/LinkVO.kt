package com.liux.blog.bean.vo

import com.liux.blog.bean.po.Link

data class LinkVO(
    var title: String,
    var url: String,
) {
    companion object {
        fun of(link: Link): LinkVO {
            return LinkVO(
                link.title!!,
                link.url!!,
            )
        }
    }
}
