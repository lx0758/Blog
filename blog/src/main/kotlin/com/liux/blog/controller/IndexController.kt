package com.liux.blog.controller

import com.liux.blog.annotation.RequestUrl
import com.liux.blog.bean.po.ArticleUrlStatusEnum
import com.liux.blog.bean.vo.*
import com.liux.blog.service.*
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.client.HttpClientErrorException

@Controller
class IndexController {

    @Autowired
    private lateinit var themeService: ThemeService
    @Autowired
    private lateinit var articleService: ArticleService
    @Autowired
    private lateinit var articleUrlService: ArticleUrlService
    @Autowired
    private lateinit var categoryService: CategoryService
    @Autowired
    private lateinit var tagService: TagService

    @GetMapping("/")
    fun index(model: Model): String {
        return page(model, 1)
    }

    @GetMapping("/page/{pageNum}")
    fun page(model: Model, @PathVariable("pageNum") pageNum: Int): String {
        val articlePage = articleService.listByPage(pageNum)
        val articles = articlePage.map { ArticlePageVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model.asMap(), "page")
    }

    @GetMapping("/archive/")
    fun archive(model: Model): String {
        return archive(model, 1)
    }

    @GetMapping("/archive/{pageNum}")
    fun archive(model: Model, @PathVariable("pageNum") pageNum: Int): String {
        val articlePage = articleService.listByArchive(pageNum)
        val articles = articlePage.map { ArticleArchiveVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model.asMap(), "archive")
    }

    @GetMapping("/category/")
    fun category(model: Model): String {
        val categorys = categoryService.listByCategory().map { CategoryVO.of(it) }
        model.addAttribute("categorys", categorys)
        return themeService.render(model.asMap(), "categorys")
    }

    @GetMapping("/category/{categoryName}")
    fun category(model: Model, @PathVariable("categoryName") categoryName: String): String {
        return category(model, categoryName, 1)
    }

    @GetMapping("/category/{categoryName}/{pageNum}")
    fun category(model: Model, @PathVariable("categoryName") categoryName: String, @PathVariable("pageNum") pageNum: Int): String {
        val category = categoryService.getByByName(categoryName) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "分类不存在")
        val articlePage = articleService.listByCategory(category.id!!, pageNum)
        val articles = articlePage.map { ArticleArchiveVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("category", category.name)
        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model.asMap(), "category")
    }

    @GetMapping("/tag/")
    fun tag(model: Model): String {
        val tags = tagService.listByCount().map { TagVO.of(it) }
        model.addAttribute("tags", tags)
        return themeService.render(model.asMap(), "tags")
    }

    @GetMapping("/tag/{tagName}")
    fun tag(model: Model, @PathVariable("tagName") tagName: String): String {
        return tag(model, tagName, 1)
    }

    @GetMapping("/tag/{tagName}/{pageNum}")
    fun tag(model: Model, @PathVariable("tagName") tagName: String, @PathVariable("pageNum") pageNum: Int): String {
        val tag = tagService.getByName(tagName) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "标签不存在")
        val articlePage = articleService.listByTag(tag.id!!, pageNum)
        val articles = articlePage.map { ArticleArchiveVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("tag", tag.name)
        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model.asMap(), "tag")
    }

    @GetMapping("/article/{path}")
    fun article(model: Model, @RequestUrl url: String, @PathVariable("path") path: String): String {
        val pathId = path.toIntOrNull()
        val articleId = if (pathId != null) {
            val articleUrl = articleUrlService.getCurrentUrlByArticleId(pathId)
            if (articleUrl != null) {
                // 如果文章有URL则重定向
                return "redirect:/article/${articleUrl.url}"
            }
            pathId
        } else {
            val articleUrl = articleUrlService.getUrlByPath(path) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "文章不存在")
            if (articleUrl.status != ArticleUrlStatusEnum.CURRENT.value) {
                // 如果使用旧的URL访问则重定向
                val articleId = articleUrl.articleId!!
                val redirectPath = articleUrlService.getCurrentUrlByArticleId(articleId)?.url ?: articleId
                return "redirect:/article/${redirectPath}"
            }
            articleUrl.articleId!!
        }

        val article = articleService.getByBlog(articleId) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "文章不存在")
        val articleVO = ArticleVO.of(article)
        val prev = articleService.getByPrev(article.id!!)
        val next = articleService.getByNext(article.id!!)
        val prevVO = if (prev != null) ArticleArchiveVO.of(prev) else null
        val nextVO = if (next != null) ArticleArchiveVO.of(next) else null

        model.addAttribute("article", articleVO)
        model.addAttribute("catalogues", articleVO.catalogues)
        model.addAttribute("prev", prevVO)
        model.addAttribute("next", nextVO)
        return themeService.render(model.asMap(), "article", articleVO.title, articleVO.description, articleVO.tags.joinToString(","), url)
    }
}
