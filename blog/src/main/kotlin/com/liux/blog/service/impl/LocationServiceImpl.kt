package com.liux.blog.service.impl

import com.liux.blog.service.LocationService
import org.lionsoul.ip2region.xdb.Searcher
import org.springframework.core.io.ClassPathResource
import org.springframework.stereotype.Service

@Service
class LocationServiceImpl : LocationService {

    private val searcher by lazy {
        val dbBytes = ClassPathResource("ip2region.xdb").inputStream.readAllBytes()
        Searcher.newWithBuffer(dbBytes)
    }

    override fun getLocationFromIp(ip: String): String? {
        try {
            val searchResult = searcher.search(ip)

            if (searchResult.isNullOrBlank()) {
                return null
            }

            return searchResult
                .split("\\|".toRegex())
                .filter { it != "0" }
                .joinToString(separator = ",")
        } catch (e: Exception) {
            return null
        }
    }
}