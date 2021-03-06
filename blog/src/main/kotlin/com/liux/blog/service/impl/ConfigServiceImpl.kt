package com.liux.blog.service.impl

import com.liux.blog.dao.ConfigMapper
import com.liux.blog.service.ConfigService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service

@Service
class ConfigServiceImpl: ConfigService {

    @Autowired
    private lateinit var configMapper: ConfigMapper

    override fun listConfigs(): Map<String, String?> {
        val configMap = HashMap<String, String?>()
        val configs = configMapper.select()
        for (config in configs) {
            configMap[config.key] = config.value
        }
        return configMap
    }
}
