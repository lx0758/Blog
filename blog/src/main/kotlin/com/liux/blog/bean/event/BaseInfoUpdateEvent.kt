package com.liux.blog.bean.event

import org.springframework.context.ApplicationEvent

class BaseInfoUpdateEvent private constructor(
    val source: String,
) : ApplicationEvent(source) {
    companion object {
        val CONFIG = BaseInfoUpdateEvent("CONFIG")
        val USER = BaseInfoUpdateEvent("USER")

        val ARTICLE = BaseInfoUpdateEvent("ARTICLE")
        val CATEGORY = BaseInfoUpdateEvent("CATEGORY")
        val TAG = BaseInfoUpdateEvent("TAG")
        val LINK = BaseInfoUpdateEvent("LINK")
    }
}