package com.liux.blog.service.impl

import com.liux.blog.bean.vo.BlogVO
import com.liux.blog.dao.ArticleMapper
import com.liux.blog.dao.CategoryMapper
import com.liux.blog.dao.TagMapper
import com.liux.blog.dao.UserMapper
import com.liux.blog.service.ConfigService
import com.liux.blog.service.LinkService
import com.liux.blog.service.ThemeService
import org.apache.commons.logging.LogFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class ThemeServiceImpl: ThemeService {

    private val logger = LogFactory.getLog(javaClass)

    @Autowired
    private lateinit var configService: ConfigService
    @Autowired
    private lateinit var articleMapper: ArticleMapper
    @Autowired
    private lateinit var categoryMapper: CategoryMapper
    @Autowired
    private lateinit var tagMapper: TagMapper
    @Autowired
    private lateinit var userMapper: UserMapper
    @Autowired
    private lateinit var linkService: LinkService

    private var cacheBlogVO: BlogVO? = null

    override fun updateBlogInfo() {
        logger.info("Update blog info, current blog info is ${if (cacheBlogVO != null) "not null" else "null"}")
        val configs = configService.listByTheme()
        val articleCount = articleMapper.getCount()
        val categoryCount = categoryMapper.getCount()
        val tagCount = tagMapper.selectCount()
        val user = userMapper.getByOwner()
        val links = linkService.listByBlog()
        cacheBlogVO = BlogVO.of(configs, articleCount, categoryCount, tagCount, user, links)
    }

    override fun getCacheBlogInfo(): BlogVO {
        if (cacheBlogVO == null) updateBlogInfo()
        return cacheBlogVO!!
    }

    override fun render(
        model: MutableMap<String, Any?>,
        page: String,
        pageTitle: String?,
        pageDescription: String?,
        pageKeywords: String?,
        pageUrl: String?,
    ): String {
        model["page"] = page
        model["pageTitle"] = pageTitle
        model["pageDescription"] = pageDescription
        model["pageKeywords"] = pageKeywords
        model["pageUrl"] = pageUrl
        model["blog"] = getCacheBlogInfo()

        if (model["pageTitle"] == null) {
            model["pageTitle"] = when (page) {
                "archive" -> "归档"
                "article" -> "文章"
                "category" -> "分类详情"
                "categorys" -> "分类"
                "page" -> null
                "tag" -> "标签详情"
                "tags" -> "标签"
                else -> null
            }
        }

        return "index"
    }
}
