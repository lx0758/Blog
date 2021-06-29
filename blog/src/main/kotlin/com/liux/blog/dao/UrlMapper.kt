package com.liux.blog.dao

import com.liux.blog.bean.po.Url
import org.springframework.stereotype.Repository

@Repository
interface UrlMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: Url): Int
    fun insertSelective(record: Url?): Int
    fun selectByKey(key: String): Url?
    fun selectByPrimaryKey(id: Int): Url?
    fun selectByAdmin(url: Url): List<Url>
    fun updateByPrimaryKeySelective(record: Url): Int
    fun updateByPrimaryKey(record: Url): Int
}
