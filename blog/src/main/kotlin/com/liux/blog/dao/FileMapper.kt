package com.liux.blog.dao

import com.liux.blog.bean.po.File
import org.springframework.stereotype.Repository

@Repository
interface FileMapper {
    fun deleteByPrimaryKey(id: Int?): Int
    fun insert(record: File?): Int
    fun insertSelective(record: File?): Int
    fun getByPrimaryKey(id: Int?): File?
    fun getCount(): Int
    fun selectByAdminDashboard(): List<File>
    fun selectByAdmin(file: File): List<File>
    fun updateByPrimaryKeySelective(file: File?): Int
    fun updateByPrimaryKey(file: File?): Int
}