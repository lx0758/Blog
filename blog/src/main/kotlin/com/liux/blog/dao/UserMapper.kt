package com.liux.blog.dao

import com.liux.blog.bean.po.User
import org.springframework.stereotype.Repository

@Repository
interface UserMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: User): Int
    fun insertSelective(record: User): Int
    fun getByPrimaryKey(id: Int): User?
    fun getByAuthor(id: Int): User?
    fun getByOwner(): User
    fun getByUsername(username: String): User?
    fun selectByAdmin(user: User): List<User>
    fun updateByPrimaryKeySelective(record: User): Int
    fun updateByPrimaryKey(record: User): Int
}
