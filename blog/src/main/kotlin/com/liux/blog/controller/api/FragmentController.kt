package com.liux.blog.controller.api

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.Resp
import com.liux.blog.bean.po.FragmentStatusEnum
import com.liux.blog.bean.po.isValid
import com.liux.blog.bean.vo.api.FragmentItemVO
import com.liux.blog.bean.vo.api.FragmentVO
import com.liux.blog.bean.vo.api.PaginationListVO
import com.liux.blog.service.FragmentService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.*
import org.springframework.web.client.HttpClientErrorException

@RestController
@RequestMapping("/api/fragment")
class FragmentController {

    @Autowired
    private lateinit var fragmentService: FragmentService

    @GetMapping
    fun query(
        @RequestParam("key", required = false) key: String?,
        @RequestParam("content", required = false) content: String?,
        @RequestParam("status", required = false) status: Int?,
        @RequestParam("pageNum") pageNum: Int,
        @RequestParam("pageSize") pageSize: Int,
        @RequestParam("orderName", required = false) orderName: String?,
        @RequestParam("orderMethod", required = false) orderMethod: String?,
    ): Resp<PaginationListVO<FragmentItemVO>> {
        val fragments = fragmentService.listByAdmin(key, content, status, pageNum, pageSize, orderName, orderMethod)
        val fragmentVOs = fragments.map{ FragmentItemVO.of(it) }
        return Resp.succeed(PaginationListVO.of(fragmentVOs, fragments))
    }

    @GetMapping("{id}")
    fun queryInfo(
        @PathVariable("id") id: Int,
    ): Resp<FragmentVO> {
        val fragment = fragmentService.getByAdmin(id) ?: throw HttpClientErrorException(HttpStatus.NOT_FOUND, "片段不存在")
        val fragmentVO = FragmentVO.of(fragment)
        return Resp.succeed(fragmentVO)
    }

    @PostMapping
    fun add(
        @CurrentUserId userId: Int,
        @RequestParam("key") key: String?,
        @RequestParam("content") content: String,
        @RequestParam("status") status: Int,
    ): Resp<FragmentVO> {
        checkParams(content, status)
        val fragment = fragmentService.addByAdmin(userId, key, content, status)
        val fragmentVO = FragmentVO.of(fragment)
        return Resp.succeed(fragmentVO)
    }

    @PutMapping("{id}")
    fun update(
        @PathVariable("id") id: Int,
        @RequestParam("key") key: String?,
        @RequestParam("content") content: String,
        @RequestParam("status") status: Int,
    ): Resp<*> {
        checkParams(content, status)
        val updateRow = fragmentService.updateByAdmin(id, key, content, status)
        if (updateRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "更新失败")
        }
        return Resp.succeed()
    }

    @DeleteMapping("{id}")
    fun delete(
        @PathVariable("id") id: Int,
    ): Resp<*> {
        val deleteRow = fragmentService.deleteByAdmin(id)
        if (deleteRow < 1) {
            throw HttpClientErrorException(HttpStatus.NOT_FOUND, "删除失败")
        }
        return Resp.succeed()
    }

    private fun checkParams(content: String, status: Int) {
        if (content.isEmpty()) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "内容不能为空")
        }
        if (!FragmentStatusEnum.entries.toTypedArray().isValid(status)) {
            throw HttpClientErrorException(HttpStatus.BAD_REQUEST, "状态不正确")
        }
    }
}
