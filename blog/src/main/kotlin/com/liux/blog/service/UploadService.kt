package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Upload

interface UploadService {
    fun listByAdmin(pageNum: Int, pageSize: Int): Page<Upload>
}