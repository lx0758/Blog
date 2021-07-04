package com.liux.blog.dao

import com.liux.blog.bean.po.Article
import org.apache.ibatis.annotations.Param
import org.springframework.stereotype.Repository

@Repository
interface ArticleMapper {

    fun deleteByPrimaryKey(id: Int): Int
    fun insert(record: Article): Int
    fun insertSelective(record: Article): Int
    fun getByPrimaryKey(id: Int): Article?
    fun getByBlog(@Param("id") id: Int?): Article?
    fun getByIdOrUrl(@Param("id") id: Int, @Param("url") url: String): Article?
    fun getByPrev(articleId: Int): Article?
    fun getByNext(articleId: Int): Article?
    fun getByAdmin(@Param("id") id: Int?): Article?
    fun getCount(): Int
    fun selectByPage(): List<Article>
    fun selectByArchive(): List<Article>
    fun selectByCategory(categoryId: Int): List<Article>
    fun selectByTag(tagId: Int): List<Article>
    fun selectBySearch(): List<Article>
    fun selectBySitemap(): List<Article>
    fun selectByAdmin(article: Article): List<Article>
    fun updateByPrimaryKeySelective(record: Article): Int
    fun updateByPrimaryKey(record: Article): Int
    fun updateByMoveCategory(@Param("oldCategoryId") oldCategoryId: Int, @Param("newCategoryId") newCategoryId: Int)
}
