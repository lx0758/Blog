package com.liux.blog.dao

import com.liux.blog.bean.po.Config
import org.springframework.stereotype.Repository

@Repository
interface ConfigMapper {
    fun deleteByPrimaryKey(key: String): Int
    fun insert(record: Config): Int
    fun insertSelective(record: Config): Int
    fun selectByPrimaryKey(key: String): Config?
    fun select(): List<Config>
    fun updateByPrimaryKeySelective(record: Config): Int
    fun updateByPrimaryKey(record: Config): Int
}
