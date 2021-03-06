package com.liux.blog

import org.junit.jupiter.api.Test
import java.io.File
import java.io.FileOutputStream
import java.sql.DriverManager
import javax.swing.filechooser.FileSystemView


class ArticleTest {

    @Test
    fun exportToHexoPage() {
        val fileSystemView = FileSystemView.getFileSystemView()
        val dir = File(fileSystemView.homeDirectory.absolutePath + "/hexo/source/_posts")

        Class.forName("com.mysql.cj.jdbc.Driver");
        val connection = DriverManager.getConnection(
            "jdbc:mysql://localhost:3306/blog",
            "root",
            "root"
        )
        val statement = connection.createStatement()
        val resultSet = statement.executeQuery("""
            SELECT
            t_article.id,
            t_article.title,
            t_article.create_time,
            t_category.`name` AS category,
            (SELECT GROUP_CONCAT(t_tag.`name`) FROM t_article_tag LEFT JOIN t_tag ON t_article_tag.tag_id = t_tag.id WHERE t_article_tag.article_id = t_article.id) AS tags,
            t_article.content
            FROM
            t_article
            LEFT JOIN t_category ON t_article.category_id = t_category.id;
            """.trimIndent())
        while (resultSet.next()) {
            val id = resultSet.getString("id")
            val title = resultSet.getString("title")
            val date = resultSet.getString("create_time")
            val category = resultSet.getString("category")
            val tags = resultSet.getString("tags")
            val content = resultSet.getString("content")

            writeMarkdown(dir, id, title, date, category, tags, content)
        }
        resultSet.close()
        statement.close()
        connection.close()
    }

    private fun writeMarkdown(
        dir: File,
        id: String,
        title: String,
        date: String,
        category: String,
        tags: String?,
        content: String
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
