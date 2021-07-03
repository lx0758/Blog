package com.liux.blog.config

import org.springframework.context.annotation.Configuration
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer
import java.io.File

@Configuration
class UploadConfig: WebMvcConfigurer {

    companion object {
        const val UPLOAD_URL_PREFIX = "/upload/"
        val UPLOAD_DIR = File(System.getProperty("user.dir"), "upload")
    }

    override fun addResourceHandlers(registry: ResourceHandlerRegistry) {
        registry
            .addResourceHandler("$UPLOAD_URL_PREFIX**")
            .addResourceLocations("file:${UPLOAD_DIR.canonicalPath}${File.separator}")
    }
}