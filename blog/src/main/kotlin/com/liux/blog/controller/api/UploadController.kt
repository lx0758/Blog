package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.bean.vo.api.UploadVO
import com.liux.blog.service.UploadService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/upload")
class UploadController {

    @Autowired
    private lateinit var uploadService: UploadService

    @GetMapping
    fun query(
        @RequestParam("name", required = false) name: String?,
        @RequestParam("type", required = false) type: String?,
        @RequestParam("status", required = false) status: Int?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
    ): Resp<PaginationListVO<UploadVO>> {
        val uploads = uploadService.listByAdmin(name, type, status, pageNum, pageSize)
        val uploadVOs = uploads.map { UploadVO.of(it) }
        return Resp.succeed(PaginationListVO.of(uploadVOs, uploads))
    }
}