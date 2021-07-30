package com.liux.blog.dao

import com.liux.blog.bean.po.Features
import org.springframework.stereotype.Repository

@Repository
interface FeaturesMapper {
    fun deleteByPrimaryKey(key: String): Int
    fun insert(record: Features): Int
    fun insertSelective(record: Features): Int
    fun getByPrimaryKey(key: String): Features?
    fun updateByPrimaryKeySelective(record: Features): Int
    fun updateByPrimaryKeyNullable(record: Features): Int
    fun updateByPrimaryKey(record: Features): Int
}