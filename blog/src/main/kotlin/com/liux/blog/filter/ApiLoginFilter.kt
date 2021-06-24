package com.liux.blog.filter

import com.liux.blog.bean.Resp
import com.liux.blog.bean.UserDetailsImpl
import com.liux.blog.checkVerifyCode
import com.liux.blog.service.UserService
import com.liux.blog.toJSONString
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.security.authentication.*
import org.springframework.security.core.Authentication
import org.springframework.security.core.AuthenticationException
import org.springframework.security.web.authentication.AbstractAuthenticationProcessingFilter
import org.springframework.security.web.util.matcher.AntPathRequestMatcher
import javax.servlet.http.HttpServletRequest
import javax.servlet.http.HttpServletResponse
import kotlin.jvm.Throws

class ApiLoginFilter(
    authenticationManager: AuthenticationManager,
    private val userService: UserService
) : AbstractAuthenticationProcessingFilter(
    AntPathRequestMatcher("/api/session", "POST"),
    authenticationManager
) {

    init {
        setAuthenticationSuccessHandler { _, response, authentication ->
            userService.refreshLastLoginTime((authentication.principal as UserDetailsImpl).getId())
            response.setHeader("content-type", MediaType.APPLICATION_JSON_VALUE)
            response.writer.print(
                Resp.succeed().toJSONString()
            )
            response.writer.flush()
        }
        setAuthenticationFailureHandler { _, response, exception ->
            val reason = when(exception) {
                is BadVerifyCodeException -> "验证码不正确"
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

    @Throws(AuthenticationException::class)
    override fun attemptAuthentication(request: HttpServletRequest, response: HttpServletResponse): Authentication {
        val username = request.getParameter("username")?.trim() ?: ""
        val password = request.getParameter("password")?.trim() ?: ""
        val verifyCode = request.getParameter("verifyCode")?.trim() ?: ""

        if (!request.checkVerifyCode(verifyCode, 1)) throw BadVerifyCodeException()

        val authenticationToken = UsernamePasswordAuthenticationToken(username, password)
        authenticationToken.details = authenticationDetailsSource.buildDetails(request)

        return authenticationManager.authenticate(authenticationToken)
    }

    class BadVerifyCodeException : AuthenticationException("Incorrect verification code")
}
