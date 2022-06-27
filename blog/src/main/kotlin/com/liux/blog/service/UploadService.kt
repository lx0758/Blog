package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Upload
import org.springframework.web.multipart.MultipartFile

interface UploadService {
    fun listByAdmin(name: String?, type: String?, pageNum: Int, pageSize: Int): Page<Upload>
    fun addByAdmin(userId: Int, file: MultipartFile): Upload
    fun updateByAdmin(userId: Int, id: Int, file: MultipartFile): Int
    fun deleteByAdmin(id: Int): Int
    fun getCountByDashboard(): Int
}