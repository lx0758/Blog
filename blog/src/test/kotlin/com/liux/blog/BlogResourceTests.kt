package com.liux.blog

import org.junit.jupiter.api.Test
import java.net.URL
import java.nio.file.Files
import java.nio.file.Paths
import java.nio.file.StandardCopyOption

class BlogResourceTests {

    @Test
    fun download() {
        val content = Files.readString(
            Paths.get("src/main/resources/templates/index.html")
        )
        "/blog/cdn-resources/(.*?)[\"']".toRegex().findAll(content).forEachIndexed { index, matchResult ->
            val urlPath = matchResult.groupValues[1]
            println("download ${index + 1}: $urlPath")
            URL("https://cdn.jsdelivr.net/$urlPath").openStream().use { inputStream ->
                val target = Paths.get("src/main/resources/static/blog/cdn-resources/$urlPath")
                Files.createDirectories(target.getParent())
                Files.copy(inputStream, target, StandardCopyOption.REPLACE_EXISTING)
            }
        }
    }
}