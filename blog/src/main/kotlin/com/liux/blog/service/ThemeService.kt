package com.liux.blog.service

import com.liux.blog.bean.vo.BlogVO
import org.springframework.ui.Model

interface ThemeService {
    fun updateBlogInfo()
    fun getCacheBlogInfo(): BlogVO
    fun render(
        model: MutableMap<String, Any?>,
        page: String,
        pageTitle: String? = null,
        pageDescription: String? = null,
        pageKeywords: String? = null,
        pageUrl: String? = null,
    ): String
}
