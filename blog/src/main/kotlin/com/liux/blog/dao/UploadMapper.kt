package com.liux.blog.dao

import com.liux.blog.bean.po.Upload
import org.springframework.stereotype.Repository

@Repository
interface UploadMapper {
    fun deleteByPrimaryKey(id: Int?): Int
    fun insert(record: Upload?): Int
    fun insertSelective(record: Upload?): Int
    fun selectByPrimaryKey(id: Int?): Upload?
    fun selectByAdmin(upload: Upload): List<Upload>
    fun updateByPrimaryKeySelective(upload: Upload?): Int
    fun updateByPrimaryKey(upload: Upload?): Int
}