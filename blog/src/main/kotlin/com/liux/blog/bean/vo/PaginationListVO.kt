package com.liux.blog.bean.vo

import com.github.pagehelper.Page

class PaginationListVO<T>(
    var pageNum: Int,
    var pageSize: Int,
    var pages: Int,
    var total: Int,
    var list: List<T>,
) {
    companion object {
        fun <T> of(list: List<T>, page: Page<*>): PaginationListVO<T> {
            return PaginationListVO(
                page.pageNum,
                page.pageSize,
                page.pages,
                page.total.toInt(),
                list
            )
        }
    }
}