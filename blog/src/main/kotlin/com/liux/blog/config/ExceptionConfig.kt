package com.liux.blog.config

import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.ControllerAdvice
import org.springframework.web.bind.annotation.ExceptionHandler
import org.springframework.web.client.HttpClientErrorException
import javax.servlet.http.HttpServletResponse

@ControllerAdvice
class ExceptionConfig {

    companion object {
        private const val ignoreHttpError = false
        fun transformHttpStatus(httpStatus: HttpStatus): HttpStatus {
            return if (ignoreHttpError) {
                HttpStatus.OK
            } else {
                httpStatus
            }
        }
    }

    @ExceptionHandler(value = [HttpClientErrorException::class])
    fun handler(response: HttpServletResponse, exception: HttpClientErrorException) {
        response.sendError(exception.rawStatusCode, exception.statusText)
    }

    @ExceptionHandler(value = [Exception::class])
    fun handler(response: HttpServletResponse, exception: Exception) {
        response.sendError(HttpStatus.INTERNAL_SERVER_ERROR.value(), exception.message)
    }
}