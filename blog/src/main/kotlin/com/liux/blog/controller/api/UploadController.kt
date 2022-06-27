package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.bean.vo.api.UploadVO
import com.liux.blog.service.UploadService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException
import org.springframework.web.multipart.MultipartFile

@RestController
@RequestMapping("/api/upload")
class UploadController {

    @Autowired
    private lateinit var uploadService: UploadService

    @GetMapping
    fun query(
        @RequestParam("name", required = false) name: String?,
        @RequestParam("type", required = false) type: String?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
    ): Resp<PaginationListVO<UploadVO>> {
        val uploads = uploadService.listByAdmin(name, type, pageNum, pageSize)
        val uploadVOs = uploads.map { UploadVO.of(it) }
        return Resp.succeed(PaginationListVO.of(uploadVOs, uploads))
    }

    @PostMapping
    fun add(
        @CurrentUserId userId: Int,
        @RequestParam("file") file: MultipartFile,
    ): Resp<UploadVO> {
        val upload = uploadService.addByAdmin(userId, file)
        val uploadVO = UploadVO.of(upload)
        return Resp.succeed(uploadVO)
    }

    @PutMapping("{id}")
    fun update(
        @CurrentUserId userId: Int,
        @PathVariable("id") id: Int,
        @RequestParam("file") file: MultipartFile,
    ): Resp<*> {
        val uploadRow = uploadService.updateByAdmin(userId, id, file)
        if (uploadRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val deleteRow = uploadService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }
}