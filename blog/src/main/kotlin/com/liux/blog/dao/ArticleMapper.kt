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
    fun getByBlog(id: Int?): Article?
    fun getByPrev(articleId: Int): Article?
    fun getByNext(articleId: Int): Article?
    fun getByAdmin(id: Int?): Article?
    fun getCount(): Int
    fun getViews(): Int
    fun selectByPage(): List<Article>
    fun selectByArchive(): List<Article>
    fun selectByCategory(categoryId: Int): List<Article>
    fun selectByTag(tagId: Int): List<Article>
    fun selectBySearch(): List<Article>
    fun selectBySitemap(): List<Article>
    fun selectByExport(): List<Article>
    fun selectByAdminDashboard(): List<Article>
    fun selectByAdmin(@Param("article") article: Article, @Param("url") url: String?): List<Article>
    fun updateByPrimaryKeySelective(record: Article): Int
    fun updateByPrimaryKeyNullable(record: Article): Int
    fun updateByPrimaryKey(record: Article): Int
    fun updateByMoveCategory(@Param("oldCategoryId") oldCategoryId: Int, @Param("newCategoryId") newCategoryId: Int)
}
