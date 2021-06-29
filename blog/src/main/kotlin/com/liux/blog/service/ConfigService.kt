package com.liux.blog.service

import com.liux.blog.bean.po.Config

interface ConfigService {

    fun listByTheme(): Map<String, String?>
    fun listByAdmin(key: String?, value: String?, description: String?): List<Config>
}
