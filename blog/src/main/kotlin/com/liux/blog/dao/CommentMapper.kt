package com.liux.blog.dao

import com.liux.blog.bean.po.Comment
import org.springframework.stereotype.Repository

@Repository
interface CommentMapper {
    fun deleteByPrimaryKey(id: Int?): Int
    fun insert(record: Comment?): Int
    fun insertSelective(record: Comment?): Int
    fun selectByArticleId(articleId: Int): List<Comment>
    fun selectByPrimaryKey(id: Int?): Comment?
    fun selectCount(): Long
    fun updateByPrimaryKeySelective(record: Comment?): Int
    fun updateByPrimaryKey(record: Comment?): Int
    fun selectByLatest(size: Int): List<Comment>
}
