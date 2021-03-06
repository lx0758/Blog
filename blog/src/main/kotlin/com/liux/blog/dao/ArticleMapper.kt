package com.liux.blog.dao

import com.liux.blog.bean.po.Article
import org.apache.ibatis.annotations.Param
import org.springframework.stereotype.Repository

@Repository
interface ArticleMapper {

    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: Article): Int
    fun insertSelective(record: Article): Int
    fun selectByPrimaryKey(id: Int): Article?
    fun selectByUrl(@Param("id") id: Int, @Param("url") url: String): Article?
    fun selectByPrev(articleId: Int): Article?
    fun selectByNext(articleId: Int): Article?
    fun selectByPage(): List<Article>
    fun selectByArchive(): List<Article>
    fun selectByCategory(categoryId: Int): List<Article>
    fun selectByTag(tagId: Int): List<Article>
    fun selectBySearch(): List<Article>
    fun selectBySitemap(): List<Article>
    fun selectCount(): Int
    fun updateByPrimaryKeySelective(record: Article): Int
    fun updateByPrimaryKey(record: Article): Int
}
