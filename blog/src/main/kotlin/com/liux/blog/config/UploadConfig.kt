package com.liux.blog.config

import org.springframework.context.annotation.Configuration
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer

@Configuration
class UploadConfig: WebMvcConfigurer {

    companion object {
        val UPLOAD_DIR_ROOT = System.getProperty("user.dir") + "/upload/"
    }

    override fun addResourceHandlers(registry: ResourceHandlerRegistry) {
        registry
            .addResourceHandler("/upload/**")
            .addResourceLocations("file:$UPLOAD_DIR_ROOT")
    }
}