package com.liux.blog.dao

import com.liux.blog.bean.po.Category
import org.springframework.stereotype.Repository

@Repository
interface CategoryMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: Category): Int
    fun insertSelective(record: Category): Int
    fun getDefaultCategory(): Category?
    fun getByPrimaryKey(id: Int): Category?
    fun getByName(name: String): Category?
    fun getCount(): Int
    fun selectByCategory(): List<Category>
    fun selectByAdmin(category: Category): List<Category>
    fun updateByPrimaryKeySelective(record: Category): Int
    fun updateByPrimaryKey(record: Category): Int
}
