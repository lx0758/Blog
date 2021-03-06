package com.liux.blog.dao

import com.liux.blog.bean.po.ShortUrl
import org.springframework.stereotype.Repository

@Repository
interface ShortUrlMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: ShortUrl): Int
    fun insertSelective(record: ShortUrl?): Int
    fun selectByKey(key: String): ShortUrl?
    fun selectByPrimaryKey(id: Int): ShortUrl?
    fun updateByPrimaryKeySelective(record: ShortUrl): Int
    fun updateByPrimaryKey(record: ShortUrl): Int
}
