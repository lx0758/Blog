package com.liux.blog.dao

import com.liux.blog.bean.po.Comment
import org.springframework.stereotype.Repository

@Repository
interface CommentMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun deleteByCleanChild(parentId: Int): Int
    fun insert(record: Comment): Int
    fun insertSelective(record: Comment): Int
    fun getByPrimaryKey(id: Int): Comment?
    fun getCount(): Int
    fun selectByArticleId(articleId: Int): List<Comment>
    fun selectByLatest(size: Int): List<Comment>
    fun selectByAdminDashboard(): List<Comment>
    fun selectByAdmin(comment: Comment): List<Comment>
    fun updateByPrimaryKeySelective(record: Comment): Int
    fun updateByPrimaryKey(record: Comment): Int
}
