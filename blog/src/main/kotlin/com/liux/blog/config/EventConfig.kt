package com.liux.blog.config

import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.service.ThemeService
import org.apache.commons.logging.LogFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.context.event.ApplicationStartedEvent
import org.springframework.context.event.EventListener
import org.springframework.scheduling.annotation.Async
import org.springframework.scheduling.annotation.EnableAsync
import org.springframework.stereotype.Component

@Component
@EnableAsync
class EventConfig {

    private val logger = LogFactory.getLog(javaClass)

    @Autowired
    private lateinit var themeService: ThemeService

    @Async
    @EventListener
    fun started(event: ApplicationStartedEvent) {
        themeService.updateBlogInfo()
    }

    @Async
    @EventListener
    fun updateBaseInfo(event: BaseInfoUpdateEvent) {
        logger.info("updateBaseInfo by:" + event.source)
        themeService.updateBlogInfo()
    }
}