package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Upload
import java.util.*

class UploadVO(
    var id: Int,
    var displayName: String,
    var length: Int,
    var type: String,
    var path: String,
    var authorId: Int,
    var authorName: String,
    var createTime: Date,
    var status: Int,
) {
    companion object {
        fun of(upload: Upload): UploadVO {
            return UploadVO(
                upload.id,
                upload.displayName!!,
                upload.length!!,
                upload.type!!,
                upload.path!!,
                upload.authorId!!,
                upload.author!!.nickname!!,
                upload.createTime!!,
                upload.status!!,
            )
        }
    }
}
