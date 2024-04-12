package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.File
import com.liux.blog.config.FileConfig
import com.liux.blog.dao.FileMapper
import com.liux.blog.service.FileService
import com.liux.blog.util.PageHelperUtil
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import org.springframework.util.FileCopyUtils
import org.springframework.web.multipart.MultipartFile
import java.io.FileOutputStream
import java.text.SimpleDateFormat
import java.util.*

@Service
class FileServiceImpl: FileService {

    @Autowired
    private lateinit var fileMapper: FileMapper
    override fun listByAdminDashboard(pageNum: Int, pageSize: Int): Page<File> {
        val page = PageHelper.startPage<File>(pageNum, pageSize)
        fileMapper.selectByAdminDashboard()
        return page
    }

    override fun listByAdmin(
        name: String?,
        type: String?,
        pageNum: Int,
        pageSize: Int,
        orderName: String?,
        orderMethod: String?,
    ): Page<com.liux.blog.bean.po.File> {
        PageHelperUtil.orderBy(PageHelperUtil.Type.UPLOAD, orderName, orderMethod)
        val page = PageHelper.startPage<com.liux.blog.bean.po.File>(pageNum, pageSize)
        fileMapper.selectByAdmin(
            com.liux.blog.bean.po.File(
                displayName = name,
                type = type,
            )
        )
        return page
    }

    override fun addByAdmin(userId: Int, multipartFile: MultipartFile): File {
        val originalFilename = multipartFile.originalFilename ?: "unnamed"
        val localPath = generatorPath(originalFilename)
        val localFile = java.io.File(FileConfig.STORE_DIR, localPath)

        if (!localFile.parentFile.exists()) {
            localFile.parentFile.mkdirs()
        }
        val fileOutputStream = FileOutputStream(localFile)
        FileCopyUtils.copy(multipartFile.inputStream, fileOutputStream)

        val file = File(
            displayName = originalFilename,
            length = multipartFile.size,
            type = multipartFile.contentType,
            path = localPath,
            authorId = userId,
            createTime = Date(),
        )
        fileMapper.insertSelective(file)
        return file
    }

    override fun updateByAdmin(userId: Int, id: Int, multipartFile: MultipartFile): Int {
        val upload = fileMapper.getByPrimaryKey(id) ?: return -1

        if (multipartFile.contentType != upload.type) return 0

        val localFile = java.io.File(FileConfig.STORE_DIR, upload.path!!)

        val fileOutputStream = FileOutputStream(localFile)
        FileCopyUtils.copy(multipartFile.inputStream, fileOutputStream)

        upload.length = multipartFile.size
        upload.authorId = userId
        upload.updateTime = Date()
        return fileMapper.updateByPrimaryKeySelective(upload)
    }

    override fun deleteByAdmin(id: Int): Int {
        val upload = fileMapper.getByPrimaryKey(id) ?: return 0
        val deleteRow = fileMapper.deleteByPrimaryKey(id)
        if (deleteRow > 0) {
            var target = java.io.File(FileConfig.STORE_DIR, upload.path!!)
            while (target.delete()) {
                val parent = target.parentFile
                if (parent == FileConfig.STORE_DIR) break
                if (!parent.listFiles().isNullOrEmpty()) break
                target = parent
            }
        }
        return deleteRow
    }

    override fun getCountByDashboard(): Int {
        return fileMapper.getCount()
    }

    fun generatorPath(originalName: String): String {
        val dir = SimpleDateFormat("yyyy/MM/").format(Date()) + generateShortUuid() + "/"
        return "$dir$originalName"
    }


    var chars = arrayOf(
        "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
        "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
        "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
    )

    private fun generateShortUuid(): String {
        val stringBuilder = StringBuilder()
        val uuid = UUID.randomUUID().toString().replace("-", "")

        for (index in uuid.indices step 4) {        //分为8组
            val sub = uuid.substring(index, index + 4)  //每组4位
            val int = Integer.parseInt(sub, 16)
            stringBuilder.append(chars[int % 0x3E])
        }
        return stringBuilder.toString()
    }
}