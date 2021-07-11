package com.liux.blog.annotation

import java.lang.annotation.Documented

@Target(AnnotationTarget.VALUE_PARAMETER)
@Retention(AnnotationRetention.RUNTIME)
@Documented
annotation class CurrentUserId()
