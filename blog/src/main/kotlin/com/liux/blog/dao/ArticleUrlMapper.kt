package com.liux.blog.dao

import com.liux.blog.bean.po.ArticleUrl
import org.springframework.stereotype.Repository

@Repository
interface ArticleUrlMapper {

    fun deleteByArticleId(articleId: Int): Int

    fun deleteByPrimaryKey(url: String): Int

    fun insert(record: ArticleUrl): Int

    fun insertSelective(record: ArticleUrl): Int

    fun getCurrentUrlByArticleId(articleId: Int): ArticleUrl?

    fun getByPrimaryKey(url: String): ArticleUrl?

    fun selectByArticleId(articleId: Int): List<ArticleUrl>

    fun updateToHistoryByArticleId(articleId: Int)

    fun updateByPrimaryKeySelective(record: ArticleUrl): Int

    fun updateByPrimaryKey(record: ArticleUrl): Int
}