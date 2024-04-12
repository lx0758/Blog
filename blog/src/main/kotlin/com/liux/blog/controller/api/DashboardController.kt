package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.ArticleVO
import com.liux.blog.bean.vo.api.CommentVO
import com.liux.blog.bean.vo.api.DashboardVO
import com.liux.blog.bean.vo.api.FileVO
import com.liux.blog.service.ArticleService
import com.liux.blog.service.CommentService
import com.liux.blog.service.ThemeService
import com.liux.blog.service.FileService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/dashboard")
class DashboardController {

    @Autowired
    private lateinit var themeService: ThemeService
    @Autowired
    private lateinit var articleService: ArticleService
    @Autowired
    private lateinit var commentService: CommentService
    @Autowired
    private lateinit var fileService: FileService

    @GetMapping
    fun query(): Resp<DashboardVO> {

        val dashboardVO = DashboardVO(
            themeService.getCacheBlogInfo().siteArticleCount,
            themeService.getCacheBlogInfo().siteCategoryCount,
            themeService.getCacheBlogInfo().siteTagCount,
            fileService.getCountByDashboard(),
            commentService.getCountByDashboard(),
            articleService.getViewsByDashboard(),

            articleService.listByAdminDashboard(1, 10).map { ArticleVO.of(it) },
            commentService.listByAdminDashboard(1, 10).map { CommentVO.of(it) },
            fileService.listByAdminDashboard(1, 10).map { FileVO.of(it) },
        )
        return Resp.succeed(dashboardVO)
    }
}