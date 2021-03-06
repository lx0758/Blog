package com.liux.blog.dao

import com.liux.blog.bean.po.User
import org.springframework.stereotype.Repository

@Repository
interface UserMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: User): Int
    fun insertSelective(record: User): Int
    fun selectByPrimaryKey(id: Int): User?
    fun selectByOwner(): User
    fun selectByUsername(username: String): User?
    fun updateByPrimaryKeySelective(record: User): Int
    fun updateByPrimaryKey(record: User): Int
}
