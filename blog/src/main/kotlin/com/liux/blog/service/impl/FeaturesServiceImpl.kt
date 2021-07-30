package com.liux.blog.service.impl

import com.liux.blog.bean.po.Features
import com.liux.blog.dao.FeaturesMapper
import com.liux.blog.service.FeaturesService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Service
import java.util.*

@Service
class FeaturesServiceImpl: FeaturesService {

    @Autowired
    private lateinit var featuresMapper: FeaturesMapper

    override fun readSMTP(): Features.SMTP? {
        return read(Features.SMTP.KEY)?.get()
    }

    override fun saveSMTP(
        enable: Boolean,
        hostname: String?,
        port: Int?,
        ssl: Boolean?,
        username: String?,
        password: String?,
        fromName: String?,
        fromEmail: String?
    ): Int {
        return save(Features.SMTP.KEY, Features.SMTP(
            enable = enable,
            hostname = hostname,
            port = port,
            ssl = ssl,
            username = username,
            password = password,
            fromName = fromName,
            fromEmail = fromEmail
        ))
    }

    private fun read(key: String): Features? {
        return featuresMapper.getByPrimaryKey(key)
    }

    private fun save(key: String, value: Any?): Int {
        val features = Features(
            key = key,
            updateTime = Date()
        )
        features.set(value)
        var result = featuresMapper.updateByPrimaryKeyNullable(features)
        if (result <= 0) {
            features.createTime = Date()
            features.updateTime = null
            result = featuresMapper.insert(features)
        }
        return result
    }
}