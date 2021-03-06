package com.liux.blog.config

import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.service.ThemeService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.context.event.ApplicationStartedEvent
import org.springframework.context.event.EventListener
import org.springframework.scheduling.annotation.Async
import org.springframework.scheduling.annotation.EnableAsync
import org.springframework.stereotype.Component

@Component
@EnableAsync
class BaseInfoEventConfig {

    @Autowired
    private lateinit var themeService: ThemeService

    @Async
    @EventListener
    fun started(event: ApplicationStartedEvent) {
        themeService.updateBase()
    }

    @Async
    @EventListener
    fun update(event: BaseInfoUpdateEvent) {
        themeService.updateBase()
    }
}