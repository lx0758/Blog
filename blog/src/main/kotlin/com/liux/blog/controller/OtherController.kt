package com.liux.blog.controller

import com.liux.blog.renderOther
import com.liux.blog.service.ArticleService
import com.liux.blog.service.CaptchaService
import com.liux.blog.service.ThemeService
import com.liux.blog.toDateString
import jakarta.servlet.http.HttpServletRequest
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.MediaType
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController
import java.util.*

@RestController
class OtherController {

    @Autowired
    private lateinit var articleService: ArticleService
    @Autowired
    private lateinit var themeService: ThemeService
    @Autowired
    private lateinit var captchaService: CaptchaService

    @GetMapping("/captcha", produces = [MediaType.IMAGE_JPEG_VALUE])
    fun captcha(request: HttpServletRequest, type: Int): ByteArray {
        return captchaService.generate(request.session, type)
    }

    @GetMapping("/search.json", produces = [MediaType.APPLICATION_JSON_VALUE])
    fun search(): List<Map<String, String>> {
        val articles = articleService.listBySearch()
        return articles.map { article ->
            HashMap<String, String>().apply {
                put("title", article.title!!)
                put("content", article.renderOther())
                put("url", "/article/" + (article.url ?: article.id.toString()))
            }
        }
    }

    @GetMapping("/robots.txt", produces = [MediaType.TEXT_PLAIN_VALUE])
    fun robots(): String {
        val domain = themeService.getCacheBlogInfo().siteDomain
        return """
            User-agent: *
            Allow: /article/
            Sitemap: https://$domain/sitemap.xml
            """.trimIndent()
    }

    @GetMapping("/sitemap.xml", produces = [MediaType.APPLICATION_XML_VALUE])
    fun sitemap(): String {
        val domain = themeService.getCacheBlogInfo().siteDomain
        val updateTime = themeService.getCacheBlogInfo().updateTime
        val articles = articleService.listBySitemap()

        val urlset = StringBuilder().apply {
            append("\t<url>\n")
            append("\t\t<loc>https://$domain/</loc>\n")
            append("\t\t<lastmod>${updateTime.toDateString()}</lastmod>\n")
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
