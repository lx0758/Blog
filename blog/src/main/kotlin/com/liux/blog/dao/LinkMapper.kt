package com.liux.blog.dao

import com.liux.blog.bean.po.Link
import org.springframework.stereotype.Repository

@Repository
interface LinkMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: Link): Int
    fun insertSelective(record: Link): Int
    fun selectByPrimaryKey(id: Int): Link?
    fun select(): List<Link>
    fun updateByPrimaryKeySelective(record: Link): Int
    fun updateByPrimaryKey(record: Link): Int
}
