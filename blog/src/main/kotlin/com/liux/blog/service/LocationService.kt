package com.liux.blog.service

interface LocationService {

    fun getLocationFromIp(ip: String) : String?
}