package com.liux.blog.bean.vo

import com.github.pagehelper.Page

data class PaginationVO(
    var pageNum: Int,
    var pageSize: Int,
    var pages: Int,
    var total: Int,
) {
    companion object {
        fun of(page: Page<*>): PaginationVO {
            return PaginationVO(
                page.pageNum,
                page.pageSize,
                page.pages,
                page.total.toInt(),
            )
        }
    }
}
