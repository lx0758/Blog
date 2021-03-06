package com.liux.blog.service

import com.liux.blog.bean.vo.BaseVO
import org.springframework.ui.Model

interface ThemeService {
    fun updateBase()
    fun getCacheBase(): BaseVO
    fun render(model: Model, template: String): String
    fun renderArticle(model: Model, title: String, keywords: String): String
}
