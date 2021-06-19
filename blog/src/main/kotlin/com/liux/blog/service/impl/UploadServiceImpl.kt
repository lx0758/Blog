package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Upload
import com.liux.blog.dao.UploadMapper
import com.liux.blog.service.UploadService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class UploadServiceImpl: UploadService {

    @Autowired
    private lateinit var uploadMapper: UploadMapper

    override fun listByAdmin(pageNum: Int, pageSize: Int): Page<Upload> {
        val page = PageHelper.startPage<Upload>(pageNum, pageSize)
        uploadMapper.selectByAdmin()
        return page
    }
}