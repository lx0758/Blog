package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Fragment
import java.util.*

data class FragmentVO(
    var id: Int,
    var key: String?,
    var content: String,
    var authorId: Int,
    var createTime: Date,
    var updateTime: Date?,
    var status: Int,
) {
    companion object {
        fun of(fragment: Fragment): FragmentVO {
            return FragmentVO(
                fragment.id!!,
                fragment.key,
                fragment.content!!,
                fragment.authorId!!,
                fragment.createTime!!,
                fragment.updateTime,
                fragment.status!!,
            )
        }
    }
}