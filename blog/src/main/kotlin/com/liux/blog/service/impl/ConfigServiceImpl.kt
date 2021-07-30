package com.liux.blog.service.impl

import com.github.pagehelper.Page
import com.github.pagehelper.PageHelper
import com.liux.blog.bean.event.BaseInfoUpdateEvent
import com.liux.blog.bean.po.Config
import com.liux.blog.dao.ConfigMapper
import com.liux.blog.service.ConfigService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.ApplicationEventPublisher
import org.springframework.stereotype.Service
import java.util.*
import kotlin.collections.HashMap

@Service
class ConfigServiceImpl: ConfigService {

    @Autowired
    private lateinit var applicationEventPublisher: ApplicationEventPublisher

    @Autowired
    private lateinit var configMapper: ConfigMapper

    override fun listByTheme(): Map<String, String?> {
        val configMap = HashMap<String, String?>()
        val configs = configMapper.selectByTheme()
        for (config in configs) {
            configMap[config.key!!] = config.value
        }
        return configMap
    }

    override fun listByAdmin(key: String?, value: String?, description: String?, pageNum: Int, pageSize: Int): Page<Config> {
        val page = PageHelper.startPage<Config>(pageNum, pageSize)
        configMapper.selectByAdmin(Config(
            key = key,
            value = value,
            description = description,
        ))
        return page
    }

    override fun addByAdmin(key: String, value: String?, description: String) {
        configMapper.insertSelective(Config(
            key = key,
            value = value,
            description = description,
            createTime = Date(),
            updateTime = null,
        ))
        applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.CONFIG)
    }

    override fun updateByAdmin(key: String, value: String?, description: String): Int {
        val updateRow = configMapper.updateByPrimaryKeyNullable(Config(
            key = key,
            value = value,
            description = description,
            updateTime = Date(),
        ))
        if (updateRow > 0) {
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.CONFIG)
        }
        return updateRow
    }

    override fun deleteByAdmin(key: String): Int {
        val deleteRow = configMapper.deleteByPrimaryKey(key)
        if (deleteRow > 0) {
            applicationEventPublisher.publishEvent(BaseInfoUpdateEvent.CONFIG)
        }
        return deleteRow
    }
}
