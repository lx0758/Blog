package com.liux.blog.filter

import com.liux.blog.bean.Resp
import com.liux.blog.toJSONString
import org.springframework.core.log.LogMessage
import org.springframework.security.core.context.SecurityContextHolder
import org.springframework.security.web.authentication.logout.LogoutSuccessHandler
import org.springframework.security.web.authentication.logout.SecurityContextLogoutHandler
import org.springframework.security.web.util.matcher.AntPathRequestMatcher
import org.springframework.web.filter.GenericFilterBean
import javax.servlet.FilterChain
import javax.servlet.ServletRequest
import javax.servlet.ServletResponse
import javax.servlet.http.HttpServletRequest
import javax.servlet.http.HttpServletResponse

class ApiLogoutFilter : GenericFilterBean() {

    private val contextLogoutHandler = SecurityContextLogoutHandler()
    private val logoutSuccessHandler = LogoutSuccessHandler { _, response, _ ->
        response.setHeader("content-type", "application/json;charset=UTF-8")
        response.writer.print(
            Resp.succeed().toJSONString()
        )
        response.writer.flush()
    }

    private val logoutRequestMatcher = AntPathRequestMatcher("/admin/api/session", "DELETE")

    override fun doFilter(request: ServletRequest, response: ServletResponse, chain: FilterChain) {
        doFilter(request as HttpServletRequest, response as HttpServletResponse, chain)
    }

    private fun requiresLogout(request: HttpServletRequest?, response: HttpServletResponse?): Boolean {
        if (this.logoutRequestMatcher.matches(request)) {
            return true
        }
        if (logger.isTraceEnabled) {
            logger.trace(LogMessage.format("Did not match request to %s", this.logoutRequestMatcher))
        }
        return false
    }

    private fun doFilter(request: HttpServletRequest, response: HttpServletResponse, chain: FilterChain) {
        if (requiresLogout(request, response)) {
            val auth = SecurityContextHolder.getContext().authentication
            if (logger.isDebugEnabled) {
                logger.debug(LogMessage.format("Logging out [%s]", auth))
            }
            contextLogoutHandler.logout(request, response, auth)
            logoutSuccessHandler.onLogoutSuccess(request, response, auth)
            return
        }
        chain.doFilter(request, response)
    }
}
