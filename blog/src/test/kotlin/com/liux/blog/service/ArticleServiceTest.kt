package com.liux.blog.service

import com.liux.blog.toDateString
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import java.io.File
import java.io.FileOutputStream
import javax.swing.filechooser.FileSystemView

@SpringBootTest
class ArticleServiceTest {

    @Autowired
    private lateinit var articleService: ArticleService

    @Test
    fun exportToHexoPage() {
        val fileSystemView = FileSystemView.getFileSystemView()
        val dir = File(fileSystemView.homeDirectory.absolutePath + "/hexo/source/_posts").apply {
            mkdirs()
        }
        val articles = articleService.listByExport()
        articles.forEach { article ->
            val id = article.id.toString()
            val title = article.title!!
            val date = article.createTime!!.toDateString()
            val category = article.category!!.name!!
            val tags = article.tags?.joinToString(separator = ",")
            val content = article.content!!

            writeMarkdown(dir, id, title, date, category, tags, content)
        }
    }

    private fun writeMarkdown(
        dir: File, id: String, title: String, date: String, category: String, tags: String?, content: String
    ) {
        val tagString = StringBuilder()
        tags?.apply {
            tagString.append("\ntags:")
            split(",").forEach {
                tagString.append("\n\t- $it")
            }
        }

        val string = """
            ---
            title: $title
            date: $date
            type: "post"
            categories: 
                - $category$tagString
            ---
            $content
            """.trimIndent()

        FileOutputStream(File(dir, "$id.md")).use {
            it.write(string.toByteArray())
            it.flush()
        }
    }
}
