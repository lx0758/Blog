package com.liux.blog.dao

import com.liux.blog.bean.po.Tag
import org.apache.ibatis.annotations.Param
import org.springframework.stereotype.Repository

@Repository
interface TagMapper {
    fun deleteByPrimaryKey(id: Int): Int
    fun deleteByCleanTagLink(tagId: Int): Int
    fun insert(record: Tag): Int
    fun insertSelective(record: Tag): Int
    fun getByPrimaryKey(id: Int): Tag?
    fun getByName(name: String): Tag?
    fun selectByArticleId(articleId: Int): List<Tag>
    fun selectByCount(): List<Tag>
    fun selectByAdmin(tag: Tag): List<Tag>
    fun getCount(): Int
    fun updateByPrimaryKeySelective(record: Tag): Int
    fun updateByPrimaryKey(record: Tag): Int

    fun getByTagCount(tagId: Int): Int
    fun insertByAddArticleLink(@Param("articleId") articleId: Int, @Param("tagId") tagId: Int)
    fun deleteByRemoveArticleLink(@Param("articleId") articleId: Int, @Param("tagId") tagId: Int)
}
