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
        uploadMapper.selectByAdmin(Upload(
            displayName = name,
            type = type,
        ))
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

    override fun deleteByAdmin(id: Int): Int {
        val upload = uploadMapper.getByPrimaryKey(id) ?: return 0
        val deleteRow = uploadMapper.deleteByPrimaryKey(id)
        if (deleteRow > 0) {
            File(UploadConfig.UPLOAD_DIR, upload.path!!).delete()
        }
        return deleteRow
    }

    private fun generatorPath(originalName: String): String {
        val suffix = originalName.split(".").last()
        val dir = SimpleDateFormat("yyyy/MM/").format(Date())
        var name = UUID.randomUUID().toString().replace("-", "")
        if (suffix.isNotEmpty()) {
            name += (".$suffix")
        }
        return "$dir$name"
    }
}