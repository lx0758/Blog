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
import org.springframework.ui.Model

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

    override fun updateBase() {
        logger.info("Update base info, current base info is ${if (cacheBlogVO != null) "not null" else "null"}")
        val configs = configService.listByTheme()
        val articleCount = articleMapper.getCount()
        val categoryCount = categoryMapper.getCount()
        val tagCount = tagMapper.selectCount()
        val user = userMapper.getByOwner()
        val links = linkService.listByBlog()
        cacheBlogVO = BlogVO.of(configs, articleCount, categoryCount, tagCount, user, links)
    }

    override fun getCacheBase(): BlogVO {
        if (cacheBlogVO == null) updateBase()
        return cacheBlogVO!!
    }

    override fun render(model: Model, template: String): String {
        if (template != "article") {
            val title = when (template) {
                "archive" -> "归档"
                "category" -> "分类详情"
                "categorys" -> "分类"
                //"page" -> "首页"
                "tag" -> "标签详情"
                "tags" -> "标签"
                else -> null
            }
            model.addAttribute("title", title)
        }
        model.addAttribute("base", getCacheBase())
        model.addAttribute("template", template)
        return "index"
    }

    override fun renderArticle(model: Model, title: String, keywords: String): String {
        model.addAttribute("title", title)
        model.addAttribute("keywords", keywords)
        return render(model, "article")
    }
}
