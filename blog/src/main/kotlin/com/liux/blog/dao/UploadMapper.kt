package com.liux.blog.dao

import com.liux.blog.bean.po.Upload
import org.springframework.stereotype.Repository

@Repository
interface UploadMapper {
    fun deleteByPrimaryKey(id: Int?): Int
    fun insert(record: Upload?): Int
    fun insertSelective(record: Upload?): Int
    fun selectByPrimaryKey(id: Int?): Upload?
    fun selectByAdmin(): List<Upload>
    fun updateByPrimaryKeySelective(record: Upload?): Int
    fun updateByPrimaryKey(record: Upload?): Int
}