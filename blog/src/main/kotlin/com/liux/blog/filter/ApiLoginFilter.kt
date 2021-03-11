package com.liux.blog.filter

import com.liux.blog.bean.Resp
import com.liux.blog.bean.UserDetailsWapper
import com.liux.blog.service.UserService
import com.liux.blog.toJSONString
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
    AntPathRequestMatcher("/admin/api/session", "POST"),
    authenticationManager
) {

    init {
        setAuthenticationSuccessHandler { _, response, authentication ->
            userService.refreshLastLoginTime((authentication.principal as UserDetailsWapper).user.id)
            response.setHeader("content-type", "application/json;charset=UTF-8")
            response.writer.print(
                Resp.succeed().toJSONString()
            )
            response.writer.flush()
        }
        setAuthenticationFailureHandler { _, response, exception ->
            response.setHeader("content-type", "application/json;charset=UTF-8")
            response.writer.print(
                when(exception) {
                    is BadCredentialsException -> Resp.failed("账号或密码不正确").toJSONString()
                    is AccountExpiredException -> Resp.failed("账户已失效").toJSONString()
                    is CredentialsExpiredException -> Resp.failed("密码已失效").toJSONString()
                    is DisabledException -> Resp.failed("账户已被禁用").toJSONString()
                    is LockedException -> Resp.failed("账户已被锁定").toJSONString()
                    else -> Resp.failed("登录失败:${exception.message}").toJSONString()
                }
            )
            response.writer.flush()
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
