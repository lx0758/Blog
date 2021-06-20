package com.liux.blog.filter

import com.liux.blog.bean.Resp
import com.liux.blog.bean.UserDetailsWapper
import com.liux.blog.service.UserService
import com.liux.blog.toJSONString
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.security.authentication.*
import org.springframework.security.core.Authentication
import org.springframework.security.web.authentication.AbstractAuthenticationProcessingFilter
import org.springframework.security.web.util.matcher.AntPathRequestMatcher
import javax.servlet.http.HttpServletRequest
import javax.servlet.http.HttpServletResponse

class ApiLoginFilter(
    authenticationManager: AuthenticationManager,
    private val userService: UserService
) : AbstractAuthenticationProcessingFilter(
    AntPathRequestMatcher("/api/session", "POST"),
    authenticationManager
) {

    init {
        setAuthenticationSuccessHandler { _, response, authentication ->
            userService.refreshLastLoginTime((authentication.principal as UserDetailsWapper).user.id)
            response.setHeader("content-type", MediaType.APPLICATION_JSON_VALUE)
            response.writer.print(
                Resp.succeed().toJSONString()
            )
            response.writer.flush()
        }
        setAuthenticationFailureHandler { _, response, exception ->
            val reason = when(exception) {
                is BadCredentialsException -> "账号或密码不正确"
                is AccountExpiredException -> "账户已失效"
                is CredentialsExpiredException -> "密码已失效"
                is DisabledException -> "账户已被禁用"
                is LockedException -> "账户已被锁定"
                else -> "登录失败:${exception.message}"
            }
            response.sendError(HttpStatus.NOT_ACCEPTABLE.value(), reason)
        }
    }

    override fun attemptAuthentication(request: HttpServletRequest?, response: HttpServletResponse?): Authentication {
        val username = request?.getParameter("username")?.trim() ?: ""
        val password = request?.getParameter("password")?.trim() ?: ""

        val authRequest = UsernamePasswordAuthenticationToken(username, password)
        authRequest.details = authenticationDetailsSource.buildDetails(request)

        return authenticationManager.authenticate(authRequest)
    }
}
