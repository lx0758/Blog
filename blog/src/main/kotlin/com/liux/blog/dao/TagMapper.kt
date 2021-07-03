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
    fun selectByPrimaryKey(id: Int): Tag?
    fun selectByName(name: String): Tag?
    fun selectByArticle(articleId: Int): List<Tag>
    fun selectByCount(): List<Tag>
    fun selectByAdmin(tag: Tag): List<Tag>
    fun selectCount(): Int
    fun updateByPrimaryKeySelective(record: Tag): Int
    fun updateByPrimaryKey(record: Tag): Int

    fun selectByTagCount(tagId: Int): Int
    fun insertByAddArticleLink(@Param("articleId") articleId: Int, @Param("tagId") tagId: Int): Int
    fun deleteByRemoveArticleLink(@Param("articleId") articleId: Int, @Param("tagId") tagId: Int): Int
}
