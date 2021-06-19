package com.liux.blog.controller

import com.liux.blog.parseContent
import com.liux.blog.service.ArticleService
import com.liux.blog.service.ThemeService
import com.liux.blog.toDateString
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController
import java.util.*
import kotlin.collections.HashMap

@RestController
class OtherController {

    @Autowired
    private lateinit var articleService: ArticleService

    @Autowired
    private lateinit var themeService: ThemeService

    @GetMapping("/search.json")
    fun search(): List<Map<String, String>> {
        val articles = articleService.listBySearch()
        return articles.map { article ->
            HashMap<String, String>().apply {
                put("title", article.title!!)
                put("content",  article.parseContent())
                put("url", "/article/" + (article.url ?: article.id.toString()))
            }
        }
    }

    @GetMapping("/robots.txt", produces = ["plain/txt;charset=UTF-8"])
    fun robots(): String {
        val domain = themeService.getCacheBase().siteDomain
        return """
            User-agent: *
            Disallow: /admin/
            Sitemap: https://$domain/sitemap.xml
            """.trimIndent()
    }

    @GetMapping("/sitemap.xml", produces = ["application/xml;charset=UTF-8"])
    fun sitemap(): String {
        val domain = themeService.getCacheBase().siteDomain
        val articles = articleService.listBySitemap()

        val urlset = StringBuilder().apply {
            append("\t<url>\n")
            append("\t\t<loc>https://$domain/</loc>\n")
            append("\t\t<lastmod>${Date().toDateString()}</lastmod>\n")
            append("\t\t<changefreq>always</changefreq>\n")
            append("\t\t<priority>1.0</priority>\n")
            append("\t</url>\n")
            articles.forEach {
                append("\t<url>\n")
                append("\t\t<loc>https://$domain/article/${it.url ?: it.id.toString()}</loc>\n")
                append("\t\t<lastmod>${(it.updateTime ?: it.createTime!!).toDateString()}</lastmod>\n")
                append("\t\t<changefreq>monthly</changefreq>\n")
                append("\t\t<priority>0.8</priority>\n")
                append("\t</url>\n")
            }
        }.toString()
        return """
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
$urlset
</urlset>
""".trimIndent()
    }
}
