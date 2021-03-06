package com.liux.blog.config

import com.google.code.kaptcha.impl.DefaultKaptcha
import com.google.code.kaptcha.util.Config
import org.springframework.context.annotation.Bean
import org.springframework.stereotype.Component
import java.util.*

@Component
class KaptchaConfig {

    @get:Bean
    val defaultKaptcha: DefaultKaptcha
        get() {
            val defaultKaptcha = DefaultKaptcha()
            val properties = Properties()
            properties["kaptcha.border"] = "no"
            properties["kaptcha.textproducer.font.color"] = "black"
            properties["kaptcha.image.width"] = "150"
            properties["kaptcha.image.height"] = "40"
            properties["kaptcha.textproducer.font.size"] = "30"
            properties["kaptcha.session.key"] = "verifyCode"
            properties["kaptcha.textproducer.char.space"] = "5"
            defaultKaptcha.config = Config(properties)
            return defaultKaptcha
        }
}