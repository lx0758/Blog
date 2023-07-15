package com.liux.blog.bean.vo.api

import com.liux.blog.bean.po.File
import com.liux.blog.config.FileConfig
import java.util.*

class FileVO(
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
        fun of(file: File): FileVO {
            return FileVO(
                file.id!!,
                file.displayName!!,
                file.length!!,
                file.type!!,
                file.path!!,
                FileConfig.URL_PREFIX + file.path!!,
                file.authorId!!,
                file.author?.nickname ?: "-",
                file.createTime!!,
                file.updateTime,
            )
        }
    }
}
