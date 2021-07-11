package com.liux.blog.service

import com.github.pagehelper.Page
import com.liux.blog.bean.po.Config

interface ConfigService {

    fun listByTheme(): Map<String, String?>
    fun listByAdmin(key: String?, value: String?, description: String?, pageNum: Int, pageSize: Int): Page<Config>
    fun addByAdmin(key: String, value: String, description: String)
    fun updateByAdmin(key: String, value: String, description: String): Int
    fun deleteByAdmin(key: String): Int
}
