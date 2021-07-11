package com.liux.blog.config

import com.liux.blog.annotation.CurrentUserId
import com.liux.blog.bean.UserDetailsImpl
import org.springframework.context.annotation.Configuration
import org.springframework.core.MethodParameter
import org.springframework.security.core.context.SecurityContextHolder
import org.springframework.web.bind.support.WebDataBinderFactory
import org.springframework.web.context.request.NativeWebRequest
import org.springframework.web.method.support.HandlerMethodArgumentResolver
import org.springframework.web.method.support.ModelAndViewContainer
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer

@Configuration
class ParameterConfig : WebMvcConfigurer {
    override fun addArgumentResolvers(argumentResolvers: MutableList<HandlerMethodArgumentResolver>) {
        argumentResolvers.add(UserIdArgumentResolver())
    }

    class UserIdArgumentResolver : HandlerMethodArgumentResolver {
        override fun supportsParameter(parameter: MethodParameter): Boolean {
            return parameter.hasParameterAnnotation(CurrentUserId::class.java) &&
                    parameter.parameterType.isAssignableFrom(Int::class.java)
        }

        override fun resolveArgument(
            parameter: MethodParameter,
            mavContainer: ModelAndViewContainer?,
            webRequest: NativeWebRequest,
            binderFactory: WebDataBinderFactory?
        ): Any? {
            val userDetailsImpl = SecurityContextHolder.getContext().authentication.principal as UserDetailsImpl
            return userDetailsImpl.getId()
        }
    }
}