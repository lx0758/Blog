package com.liux.blog.config

import com.liux.blog.bean.Resp
import com.liux.blog.bean.UserDetailsImpl
import com.liux.blog.service.CaptchaService
import com.liux.blog.service.UserService
import com.liux.blog.toJSONString
import jakarta.servlet.http.HttpServletRequest
import jakarta.servlet.http.HttpServletResponse
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.http.HttpHeaders
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.security.authentication.*
import org.springframework.security.authentication.dao.DaoAuthenticationProvider
import org.springframework.security.config.Customizer
import org.springframework.security.config.annotation.web.builders.HttpSecurity
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity
import org.springframework.security.config.annotation.web.configurers.CsrfConfigurer
import org.springframework.security.core.Authentication
import org.springframework.security.core.AuthenticationException
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder
import org.springframework.security.crypto.password.PasswordEncoder
import org.springframework.security.web.SecurityFilterChain
import org.springframework.security.web.authentication.AbstractAuthenticationProcessingFilter
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter
import org.springframework.security.web.authentication.logout.SecurityContextLogoutHandler
import org.springframework.security.web.util.matcher.AntPathRequestMatcher
import org.springframework.security.web.util.matcher.RequestMatcher

@Configuration
@EnableWebSecurity
class SecurityConfig {

    @Autowired
    private lateinit var userService: UserService

    @Autowired
    private lateinit var captchaService: CaptchaService

    @Bean
    fun passwordEncoder(): PasswordEncoder {
        return BCryptPasswordEncoder()
    }

    @Bean
    fun authenticationProvider(passwordEncoder: PasswordEncoder): AuthenticationProvider {
        return DaoAuthenticationProvider().apply {
            setUserDetailsService(userService)
            setPasswordEncoder(passwordEncoder)
        }
    }

    @Bean
    fun authenticationManager(authenticationProvider: AuthenticationProvider): AuthenticationManager {
        return ProviderManager(authenticationProvider)
    }

    @Bean
    fun securityFilterChain(http: HttpSecurity, authenticationManager: AuthenticationManager): SecurityFilterChain {
        return http
            .cors(Customizer.withDefaults())
            .csrf { configurer -> configurer.disable() }

            .securityContext { configurer -> configurer.requireExplicitSave(false) }

            .addFilterAt(
                CaptchaLoginFilter(
                    AntPathRequestMatcher("/api/session", "POST"),
                    authenticationManager,
                    userService,
                    captchaService
                ),
                UsernamePasswordAuthenticationFilter::class.java
            )

            .logout { configurer ->
                configurer
                    .logoutRequestMatcher(AntPathRequestMatcher("/api/session", "DELETE"))
                    .clearAuthentication(true)
                    .addLogoutHandler(SecurityContextLogoutHandler())
                    .logoutSuccessHandler { _, response, _ ->
                        response.setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
                        response.writer.print(
                            Resp.succeed().toJSONString()
                        )
                        response.writer.flush()
                    }
            }

            .exceptionHandling { configurer ->
                configurer
                    .authenticationEntryPoint { _, response, _ ->
                        response.sendError(HttpStatus.UNAUTHORIZED.value(), "未登录")
                    }
                    .accessDeniedHandler { _, response, _ ->
                        response?.sendError(HttpStatus.UNAUTHORIZED.value(), "无访问权限")
                    }
            }

            .authorizeHttpRequests { configurer ->
                configurer
                    .requestMatchers("/api/captcha").permitAll()
                    .requestMatchers("/api/**").authenticated()
                    .anyRequest().permitAll()
            }

            .build()
    }

    class CaptchaLoginFilter(
        requestMatcher: RequestMatcher,
        authenticationManager: AuthenticationManager,
        private val userService: UserService,
        private val captchaService: CaptchaService
    ) : AbstractAuthenticationProcessingFilter(
        requestMatcher, authenticationManager
    ) {
        init {
            setAuthenticationSuccessHandler { _, response, authentication ->
                userService.updateByLogin((authentication.principal as UserDetailsImpl).getId())
                response.setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
                response.writer.print(
                    Resp.succeed().toJSONString()
                )
                response.writer.flush()
            }
            setAuthenticationFailureHandler { _, response, exception ->
                val reason = when (exception) {
                    is InvalidCaptchaException -> "验证码不正确"
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
            val username = request.getParameter("username")?.trim() ?: throw BadCredentialsException(null)
            val password = request.getParameter("password")?.trim() ?: throw BadCredentialsException(null)
            val captcha = request.getParameter("captcha")?.trim() ?: throw BadCredentialsException(null)

            if (!captchaService.verify(request.session, captcha, CaptchaService.TYPE_LOGIN, 1)) throw InvalidCaptchaException()

            val authenticationToken = UsernamePasswordAuthenticationToken.unauthenticated(username, password)
            authenticationToken.details = authenticationDetailsSource.buildDetails(request)
            return authenticationManager.authenticate(authenticationToken)
        }

        class InvalidCaptchaException : AuthenticationException("Incorrect captcha")
    }
}