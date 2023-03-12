package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.Upload
import com.liux.blog.config.UploadConfig
import java.util.*

class UploadVO(
    var id: Int,
    var displayName: String,
    var length: Long,
    var type: String,
    var path: String,
    var url: String,
    var authorId: Int,
    var authorName: String,
    var createTime: Date,
    var updateTime: Date?,
) {
    companion object {
        fun of(upload: Upload): UploadVO {
            return UploadVO(
                upload.id!!,
                upload.displayName!!,
                upload.length!!,
                upload.type!!,
                upload.path!!,
                UploadConfig.UPLOAD_URL_PREFIX + upload.path!!,
                upload.authorId!!,
                upload.author?.nickname ?: "-",
                upload.createTime!!,
                upload.updateTime,
            )
        }
    }
}
