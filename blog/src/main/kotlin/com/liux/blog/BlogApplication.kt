package com.liux.blog

import org.mybatis.spring.annotation.MapperScan
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
@MapperScan("com.liux.blog.dao")
class BlogApplication {
    companion object {
        val VERSION = BlogApplication::class.java.getPackage().implementationVersion
    }
}

fun main(args: Array<String>) {
    runApplication<BlogApplication>(*args)
}
