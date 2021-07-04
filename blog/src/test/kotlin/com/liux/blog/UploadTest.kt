package com.liux.blog

import org.junit.jupiter.api.Test
import org.springframework.boot.web.server.MimeMappings
import java.io.File
import java.nio.file.Files
import java.nio.file.attribute.FileTime
import java.sql.DriverManager
import java.text.SimpleDateFormat
import java.util.*
import javax.swing.filechooser.FileSystemView
import kotlin.collections.ArrayList

class UploadTest {

    @Test
    fun repairDatabase() {
        val fileSystemView = FileSystemView.getFileSystemView()
        val dir = File(fileSystemView.homeDirectory.absolutePath + "/upload/")

        val files = ArrayList<File>()
        findFile(dir, files)
        files.sortWith { o1, o2 -> (o1.getLastModifiedTime() - o2.getLastModifiedTime()).toInt() }

        Class.forName("com.mysql.cj.jdbc.Driver")
        val connection = DriverManager.getConnection(
            "jdbc:mysql://localhost:3306/blog",
            "root",
            "root"
        )
        connection.autoCommit = false
        try {
            val dirLength = dir.canonicalPath.length
            val format = SimpleDateFormat("yyyyMMddHHmmss")
            files.forEach { file ->
                val createTime = Date(file.getLastModifiedTime())
                val displayName = "微信截图_${format.format(createTime)}.${getFileExtension(file)}"
                val length = file.length()
                val type = getFileMediaType(file)
                val path = file.canonicalPath.substring(dirLength + 1).replace(File.separator, "/")

                val sql = """
                    INSERT INTO t_upload
                    (display_name, length, type, path, author_id, create_time)
                    VALUES (?, ?, ?, ?, 1, ?);
                """.trimIndent()
                val preparedStatement = connection.prepareStatement(sql)
                preparedStatement.setString(1, displayName)
                preparedStatement.setLong(2, length)
                preparedStatement.setString(3, type)
                preparedStatement.setString(4, path)
                preparedStatement.setTimestamp(5, java.sql.Timestamp(createTime.time))
                preparedStatement.execute()
                preparedStatement.close()
            }
            connection.commit()
        } catch (e: Exception) {
            connection.rollback()
            e.printStackTrace()
        }
        connection.close()
    }

    private fun findFile(file: File, files: ArrayList<File>) {
        if (file.isFile) {
            files.add(file)
            return
        }
        if (file.isDirectory) {
            file.listFiles()?.forEach {
                findFile(it, files)
            }
        }
    }

    private fun getFileExtension(file: File): String {
        return file.name.split(".").last()
    }

    private fun getFileMediaType(file: File): String {
        val extension = getFileExtension(file)
        return MimeMappings.DEFAULT.get(extension)!!
    }
}

private fun File.getLastModifiedTime(): Long {
    if (1 == 2) return lastModified()
    val attributes = Files.readAttributes(toPath(), "*")
    return (attributes["lastModifiedTime"] as FileTime).toMillis()
}
