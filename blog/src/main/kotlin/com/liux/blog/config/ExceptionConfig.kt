package com.liux.blog.config

import jakarta.servlet.http.HttpServletResponse
import org.springframework.http.HttpStatus
import org.springframework.web.bind.annotation.ControllerAdvice
import org.springframework.web.bind.annotation.ExceptionHandler
import org.springframework.web.client.HttpClientErrorException

@ControllerAdvice
class ExceptionConfig {

    companion object {
        private const val ignoreHttpError = true
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
        response.sendError(exception.statusCode.value(), exception.statusText)
    }

    @ExceptionHandler(value = [Exception::class])
    fun handler(response: HttpServletResponse, exception: Exception) {
        response.sendError(HttpStatus.INTERNAL_SERVER_ERROR.value(), exception.message)
    }
}