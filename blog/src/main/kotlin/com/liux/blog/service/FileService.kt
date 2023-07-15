package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.File
import org.springframework.web.multipart.MultipartFile

interface FileService {
    fun listByAdmin(
        name: String?,
        type: String?,
        pageNum: Int,
        pageSize: Int,
        orderName: String?,
        orderMethod: String?,
    ): Page<File>
    fun addByAdmin(userId: Int, multipartFile: MultipartFile): File
    fun updateByAdmin(userId: Int, id: Int, multipartFile: MultipartFile): Int
    fun deleteByAdmin(id: Int): Int
    fun getCountByDashboard(): Int
}