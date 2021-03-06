package com.liux.blog.dao

import com.liux.blog.bean.po.Link
import org.springframework.stereotype.Repository

@Repository
interface LinkMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: Link): Int
    fun insertSelective(record: Link): Int
    fun getByPrimaryKey(id: Int): Link?
    fun selectByBlog(): List<Link>
    fun selectByAdmin(link: Link): List<Link>
    fun updateByPrimaryKeySelective(record: Link): Int
    fun updateByPrimaryKey(record: Link): Int
}
