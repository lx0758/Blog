package com.liux.blog.config

import com.liux.blog.filter.ApiLoginFilter
import com.liux.blog.filter.ApiLogoutFilter
import com.liux.blog.service.UserService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder
import org.springframework.security.config.annotation.web.builders.HttpSecurity
import org.springframework.security.config.annotation.web.builders.WebSecurity
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter
import org.springframework.security.core.AuthenticationException
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.security.crypto.password.PasswordEncoder
import org.springframework.security.web.AuthenticationEntryPoint
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter
import org.springframework.security.web.authentication.logout.LogoutFilter
import javax.servlet.http.HttpServletRequest
import javax.servlet.http.HttpServletResponse

@Configuration
class SecurityConfig: WebSecurityConfigurerAdapter() {

    @Autowired
    private lateinit var userService: UserService

    @Bean
    fun passwordEncoder(): PasswordEncoder {
        return BCryptPasswordEncoder()
    }

    override fun configure(auth: AuthenticationManagerBuilder) {
        auth.userDetailsService(userService).passwordEncoder(passwordEncoder())
    }

    override fun configure(web: WebSecurity) {
        super.configure(web)
    }

    override fun configure(http: HttpSecurity) {
        http
            .cors()

        http
            .csrf()
            .disable()

        http
            .authorizeRequests()
            .antMatchers("/admin/api/**").authenticated()
            .anyRequest().permitAll()

        http
            .addFilterAt(ApiLoginFilter(authenticationManager(), userService), UsernamePasswordAuthenticationFilter::class.java)
            .addFilterAt(ApiLogoutFilter(), LogoutFilter::class.java)
            .exceptionHandling().authenticationEntryPoint(Http401UnauthorizedEntryPoint())
    }

    class Http401UnauthorizedEntryPoint : AuthenticationEntryPoint {
        override fun commence(
            request: HttpServletRequest,
            response: HttpServletResponse,
            exception: AuthenticationException?
        ) {
            response.sendError(HttpServletResponse.SC_UNAUTHORIZED)
        }
    }
}