package com.liux.blog.controller.blog

import com.liux.blog.exception.HttpNotFoundException
import com.liux.blog.bean.vo.*
import com.liux.blog.service.ArticleService
import com.liux.blog.service.CategoryService
import com.liux.blog.service.TagService
import com.liux.blog.service.ThemeService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable

@Controller
class IndexController {

    @Autowired
    private lateinit var themeService: ThemeService
    @Autowired
    private lateinit var articleService: ArticleService
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
        val articles = articlePage.map { ArticleItemVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model, "page")
    }

    @GetMapping("/archive/")
    fun archive(model: Model): String {
        return archive(model, 1)
    }

    @GetMapping("/archive/{pageNum}")
    fun archive(model: Model, @PathVariable("pageNum") pageNum: Int): String {
        val articlePage = articleService.listByArchive(pageNum)
        val articles = articlePage.map { ArticleSimpleVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model, "archive")
    }

    @GetMapping("/category/")
    fun category(model: Model): String {
        val categorys = categoryService.list().map { CategoryVO.of(it) }
        model.addAttribute("categorys", categorys)
        return themeService.render(model, "categorys")
    }

    @GetMapping("/category/{category}")
    fun category(model: Model, @PathVariable("category") category: String): String {
        return category(model, category, 1)
    }

    @GetMapping("/category/{category}/{pageNum}")
    fun category(model: Model, @PathVariable("category") category: String, @PathVariable("pageNum") pageNum: Int): String {
        val category = categoryService.getByName(category) ?: throw HttpNotFoundException()
        val articlePage = articleService.listByCategory(category.id, pageNum)
        val articles = articlePage.map { ArticleSimpleVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("category", category.name)
        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model, "category")
    }

    @GetMapping("/tag/")
    fun tag(model: Model): String {
        val tags = tagService.list().map { TagVO.of(it) }
        model.addAttribute("tags", tags)
        return themeService.render(model, "tags")
    }

    @GetMapping("/tag/{tag}")
    fun tag(model: Model, @PathVariable("tag") tag: String): String {
        return tag(model, tag, 1)
    }

    @GetMapping("/tag/{tag}/{pageNum}")
    fun tag(model: Model, @PathVariable("tag") tag: String, @PathVariable("pageNum") pageNum: Int): String {
        val tag = tagService.getByName(tag) ?: throw HttpNotFoundException()
        val articlePage = articleService.listByTag(tag.id, pageNum)
        val articles = articlePage.map { ArticleSimpleVO.of(it) }
        val paginationVO = PaginationVO.of(articlePage)

        model.addAttribute("tag", tag.name)
        model.addAttribute("articles", articles)
        model.addAttribute("pagination", paginationVO)
        return themeService.render(model, "tag")
    }

    @GetMapping("/article/{article}")
    fun article(model: Model, @PathVariable("article") article: String): String {
        val article = articleService.getArticleByUrl(article) ?: throw HttpNotFoundException()
        val articleVO = ArticleVO.of(article)
        val prev = articleService.getArticleByPrev(article.id)
        val next = articleService.getArticleByNext(article.id)
        val prevVO = if (prev != null) ArticleSimpleVO.of(prev) else null
        val nextVO = if (next != null) ArticleSimpleVO.of(next) else null

        model.addAttribute("article", articleVO)
        model.addAttribute("catalogues", articleVO.catalogues)
        model.addAttribute("prev", prevVO)
        model.addAttribute("next", nextVO)
        return themeService.renderArticle(model, articleVO.title, articleVO.tags.joinToString(","))
    }
}
