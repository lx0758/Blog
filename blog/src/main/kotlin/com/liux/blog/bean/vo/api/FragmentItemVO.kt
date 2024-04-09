package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Fragment
import java.util.*

data class FragmentItemVO(
    var id: Int,
    var key: String?,
    var authorId: Int,
    var authorName: String,
    var createTime: Date,
    var updateTime: Date?,
    var status: Int,
) {
    companion object {
        fun of(fragment: Fragment): FragmentItemVO {
            return FragmentItemVO(
                fragment.id!!,
                fragment.key,
                fragment.authorId!!,
                fragment.author!!.nickname!!,
                fragment.createTime!!,
                fragment.updateTime,
                fragment.status!!,
            )
        }
    }
}