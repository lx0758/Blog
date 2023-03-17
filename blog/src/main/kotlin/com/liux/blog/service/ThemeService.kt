package com.liux.blog.service

import com.liux.blog.bean.vo.BlogVO
import org.springframework.ui.Model

interface ThemeService {
    fun updateBase()
    fun getCacheBase(): BlogVO
    fun render(template: String, model: MutableMap<String, Any?>): String
    fun renderArticle(model: MutableMap<String, Any?>, title: String, keywords: String): String
}
