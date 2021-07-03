package com.liux.blog.controller.api

import com.liux.blog.bean.Resp
import com.liux.blog.bean.vo.api.DashboardVO
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/dashboard")
class DashboardController {

    @GetMapping
    fun query(): Resp<DashboardVO> {
        // TODO: 2021/7/3
        val dashboardVO = DashboardVO()
        return Resp.succeed(dashboardVO)
    }
}