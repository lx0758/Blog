package com.liux.blog.dao

import com.liux.blog.bean.po.Category
import org.springframework.stereotype.Repository

@Repository
interface CategoryMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: Category): Int
    fun insertSelective(record: Category): Int
    fun selectByPrimaryKey(id: Int): Category?
    fun selectByName(name: String): Category?
    fun selectByCategory(): List<Category>
    fun selectByAdmin(category: Category): List<Category>
    fun selectCount(): Int
    fun updateByPrimaryKeySelective(record: Category): Int
    fun updateByPrimaryKey(record: Category): Int
}
