package com.liux.blog.config

import com.google.code.kaptcha.Constants
import com.google.code.kaptcha.servlet.KaptchaServlet
import org.springframework.boot.web.servlet.ServletRegistrationBean
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

@Configuration
class KaptchaConfig {

    @Bean
    fun registerKaptchaServlet(): ServletRegistrationBean<KaptchaServlet> {
        val servletRegistrationBean = ServletRegistrationBean(KaptchaServlet(), "/vcode");
        servletRegistrationBean.initParameters.apply {
            put(Constants.KAPTCHA_SESSION_CONFIG_KEY, Constants.KAPTCHA_SESSION_KEY)
            put(Constants.KAPTCHA_SESSION_CONFIG_DATE, Constants.KAPTCHA_SESSION_DATE)
            put(Constants.KAPTCHA_BORDER, "no")
            put(Constants.KAPTCHA_BORDER_COLOR, "")
            put(Constants.KAPTCHA_BORDER_THICKNESS, "")
            put(Constants.KAPTCHA_NOISE_COLOR, "")
            put(Constants.KAPTCHA_NOISE_IMPL, "")
            put(Constants.KAPTCHA_OBSCURIFICATOR_IMPL, "")
            put(Constants.KAPTCHA_PRODUCER_IMPL, "")
            put(Constants.KAPTCHA_TEXTPRODUCER_IMPL, "")
            put(Constants.KAPTCHA_TEXTPRODUCER_CHAR_STRING, "")
            put(Constants.KAPTCHA_TEXTPRODUCER_CHAR_LENGTH, "4")
            put(Constants.KAPTCHA_TEXTPRODUCER_CHAR_SPACE, "5")
            put(Constants.KAPTCHA_TEXTPRODUCER_FONT_NAMES, "")
            put(Constants.KAPTCHA_TEXTPRODUCER_FONT_COLOR, "black")
            put(Constants.KAPTCHA_TEXTPRODUCER_FONT_SIZE, "40")
            put(Constants.KAPTCHA_WORDRENDERER_IMPL, "")
            put(Constants.KAPTCHA_BACKGROUND_IMPL, "")
            put(Constants.KAPTCHA_BACKGROUND_CLR_FROM, "")
            put(Constants.KAPTCHA_BACKGROUND_CLR_TO, "")
            put(Constants.KAPTCHA_IMAGE_WIDTH, "150")
            put(Constants.KAPTCHA_IMAGE_HEIGHT, "50")
        }
        return servletRegistrationBean;
    }
}