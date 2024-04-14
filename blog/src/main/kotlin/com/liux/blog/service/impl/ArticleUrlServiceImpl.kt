package com.liux.blog.service.impl

import com.liux.blog.bean.po.ArticleUrl
import com.liux.blog.dao.ArticleUrlMapper
import com.liux.blog.service.ArticleUrlService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class ArticleUrlServiceImpl : ArticleUrlService {

    @Autowired
    private lateinit var articleUrlMapper: ArticleUrlMapper

    override fun getUrlByPath(url: String): ArticleUrl? {
        return articleUrlMapper.getByPrimaryKey(url)
    }

    override fun getCurrentUrlByArticleId(articleId: Int): ArticleUrl? {
        return articleUrlMapper.getCurrentUrlByArticleId(articleId)
    }
}