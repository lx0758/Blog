package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.bean.vo.api.FileVO
import com.liux.blog.service.FileService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException
import org.springframework.web.multipart.MultipartFile

@RestController
@RequestMapping("/api/file")
class FileController {

    @Autowired
    private lateinit var fileService: FileService

    @GetMapping
    fun query(
        @RequestParam("name", required = false) name: String?,
        @RequestParam("type", required = false) type: String?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
        @RequestParam("orderName", required = false) orderName: String?,
        @RequestParam("orderMethod", required = false) orderMethod: String?,
    ): Resp<PaginationListVO<FileVO>> {
        val uploads = fileService.listByAdmin(name, type, pageNum, pageSize, orderName, orderMethod)
        val fileVOs = uploads.map { FileVO.of(it) }
        return Resp.succeed(PaginationListVO.of(fileVOs, uploads))
    }

    @PostMapping
    fun add(
        @CurrentUserId userId: Int,
        @RequestParam("file") file: MultipartFile,
    ): Resp<FileVO> {
        val upload = fileService.addByAdmin(userId, file)
        val fileVO = FileVO.of(upload)
        return Resp.succeed(fileVO)
    }

    @PutMapping("{id}")
    fun update(
        @CurrentUserId userId: Int,
        @PathVariable("id") id: Int,
        @RequestParam("file") file: MultipartFile,
    ): Resp<*> {
        val uploadRow = fileService.updateByAdmin(userId, id, file)
        if (uploadRow == 0) {
            throw HttpClientErrorException(HttpStatus.NOT_ACCEPTABLE, "禁止不同类型文件覆盖更新")
        }
        if (uploadRow < 0) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val deleteRow = fileService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }
}