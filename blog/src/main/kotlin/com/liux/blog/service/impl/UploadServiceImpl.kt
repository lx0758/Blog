package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.po.Upload
import com.liux.blog.config.UploadConfig
import com.liux.blog.dao.UploadMapper
import com.liux.blog.service.UploadService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import org.springframework.util.FileCopyUtils
import org.springframework.web.multipart.MultipartFile
import java.io.File
import java.io.FileOutputStream
import java.text.SimpleDateFormat
import java.util.*

@Service
class UploadServiceImpl: UploadService {

    @Autowired
    private lateinit var uploadMapper: UploadMapper

    override fun listByAdmin(name: String?, type: String?, pageNum: Int, pageSize: Int): Page<Upload> {
        val page = PageHelper.startPage<Upload>(pageNum, pageSize)
        uploadMapper.selectByAdmin(
            Upload(
                displayName = name,
                type = type,
            )
        )
        return page
    }

    override fun addByAdmin(userId: Int, file: MultipartFile): Upload {
        val originalFilename = file.originalFilename ?: "unnamed"
        val localPath = generatorPath(originalFilename)
        val localFile = File(UploadConfig.UPLOAD_DIR, localPath)

        if (!localFile.parentFile.exists()) {
            localFile.parentFile.mkdirs()
        }
        val fileOutputStream = FileOutputStream(localFile)
        FileCopyUtils.copy(file.inputStream, fileOutputStream)

        val upload = Upload(
            displayName = originalFilename,
            length = file.size,
            type = file.contentType,
            path = localPath,
            authorId = userId,
            createTime = Date(),
        )
        uploadMapper.insertSelective(upload)
        return upload
    }

    override fun updateByAdmin(userId: Int, id: Int, file: MultipartFile): Int {
        val upload = uploadMapper.getByPrimaryKey(id) ?: return 0

        if (file.contentType != upload.type) return 0

        val localFile = File(UploadConfig.UPLOAD_DIR, upload.path!!)

        val fileOutputStream = FileOutputStream(localFile)
        FileCopyUtils.copy(file.inputStream, fileOutputStream)

        upload.length = file.size
        upload.authorId = userId
        upload.updateTime = Date()
        return uploadMapper.updateByPrimaryKeySelective(upload)
    }

    override fun deleteByAdmin(id: Int): Int {
        val upload = uploadMapper.getByPrimaryKey(id) ?: return 0
        val deleteRow = uploadMapper.deleteByPrimaryKey(id)
        if (deleteRow > 0) {
            var target = File(UploadConfig.UPLOAD_DIR, upload.path!!)
            while (target.delete()) {
                val parent = target.parentFile
                if (parent == UploadConfig.UPLOAD_DIR) break
                if (!parent.listFiles().isNullOrEmpty()) break
                target = parent
            }
        }
        return deleteRow
    }

    override fun getCountByDashboard(): Int {
        return uploadMapper.getCount()
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
            val sub = uuid.substring(index, index + 4);  //每组4位
            val int = Integer.parseInt(sub, 16);
            stringBuilder.append(chars[int % 0x3E])
        }
        return stringBuilder.toString()
    }
}