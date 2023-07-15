package com.liux.blog.config

import org.springframework.context.annotation.Configuration
import org.springframework.web.servlet.config.annotation.PathMatchConfigurer
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer
import org.springframework.web.util.UrlPathHelper
import java.io.File
import java.nio.charset.StandardCharsets

@Configuration
class FileConfig: WebMvcConfigurer {

    companion object {
        const val URL_PREFIX = "/files/"
        val STORE_DIR = File(System.getProperty("user.dir"), "files")
    }

    override fun configurePathMatch(configurer: PathMatchConfigurer) {
        configurer.setUrlPathHelper(UrlPathHelper().apply {
            isUrlDecode = false
            setDefaultEncoding(StandardCharsets.UTF_8.name())
        })
        super.configurePathMatch(configurer)
    }

    override fun addResourceHandlers(registry: ResourceHandlerRegistry) {
        registry
            .addResourceHandler("$URL_PREFIX**")
            .addResourceLocations("file:${STORE_DIR.canonicalPath}${File.separator}")
    }
}