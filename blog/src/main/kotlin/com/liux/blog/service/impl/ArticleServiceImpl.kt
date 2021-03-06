package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Article
import com.liux.blog.dao.ArticleMapper
import com.liux.blog.service.ArticleService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class ArticleServiceImpl: ArticleService {

    @Autowired
    private lateinit var articleMapper: ArticleMapper

    override fun listByPage(pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByPage()
        return page
    }

    override fun listByArchive(pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByArchive()
        return page
    }

    override fun listByCategory(categoryId: Int, pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByCategory(categoryId)
        return page
    }

    override fun listByTag(tagId: Int, pageNum: Int): Page<Article> {
        val page = PageHelper.startPage<Article>(pageNum, 10)
        articleMapper.selectByTag(tagId)
        return page
    }

    override fun listBySearch(): List<Article> {
        return articleMapper.selectBySearch()
    }

    override fun listBySitemap(): List<Article> {
        return articleMapper.selectBySitemap()
    }

    override fun getArticleByUrl(url: String): Article? {
        val id = try {
            url.toInt()
        } catch (e: NumberFormatException) {
            0
        }
        return articleMapper.selectByUrl(id, url)?.apply {
            views = views?.plus(1)
            articleMapper.updateByPrimaryKeySelective(Article(id, views = views))
        }
    }

    override fun getArticleByPrev(articleId: Int): Article? {
        return articleMapper.selectByPrev(articleId)
    }

    override fun getArticleByNext(articleId: Int): Article? {
        return articleMapper.selectByNext(articleId)
    }
}
